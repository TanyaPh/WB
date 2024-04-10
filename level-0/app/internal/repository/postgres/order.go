package postgres

import (
	"api/internal/entity"
	"api/pkg/postgres"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	// "github.com/randallmlough/pgxscan"
	// pgs "github.com/georgysavva/scany/v2/pgxscan"
)

type OrderPostgres struct {
	db *postgres.Postgres
}

func NewOrderPostgres(db *postgres.Postgres) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (pg *OrderPostgres) Create(order entity.Order) error {
	ctx := context.Background()
	tx, err := pg.db.DB.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	queryDelivery, args := insertDeliveryQuery(order.Delivery)
	if err = tx.QueryRow(ctx, queryDelivery, args).Scan(&order.Delivery.Id); err != nil {
		return err
	}

	queryOrder, args := insertOrderQuery(order)
	if _, err = tx.Exec(ctx, queryOrder, args); err != nil {
		log.Println(err)
		return err
	}
	
	queryPayment, args := insertPaymentQuery(order.Payment)
	if _, err = tx.Exec(ctx, queryPayment, args); err != nil {
		return err
	}

	for i := range order.Items {
		queryItem, args := insertItemQuery(order.Items[i])
		if _, err = tx.Exec(ctx, queryItem, args); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func insertDeliveryQuery(d entity.Delivery) (string, pgx.NamedArgs) {
	query := "INSERT INTO delivery VALUES (DEFAULT, @Name, @Phone, @Zip, @City, @Address, @Region, @Email) RETURNING id"
	args := pgx.NamedArgs{
		"Name": d.Name,
		"Phone": d.Phone,
		"Zip": d.Zip,
		"City": d.City,
		"Address": d.Address,
		"Region": d.Region,
		"Email": d.Email,
	}

	return query, args
}

func insertPaymentQuery(p entity.Payment) (string, pgx.NamedArgs) {
	query := `INSERT INTO payments VALUES (@Transaction, @RequestID, @Currency, @Provider, 
						@Amount, @PaymentDT, @Bank, @DeliveryCost, @GoodTotal, @CustomFree)`

	args := pgx.NamedArgs{
		"Transaction": p.Transaction,
		"RequestID": p.RequestID,
		"Currency": p.Currency,
		"Provider": p.Provider,
		"Amount": p.Amount,
		"PaymentDT": p.PaymentDT,
		"Bank": p.Bank,
		"DeliveryCost": p.DeliveryCost,
		"GoodTotal": p.GoodTotal,
		"CustomFree": p.CustomFree,
	}

	return query, args
}

func insertOrderQuery(or entity.Order) (string, pgx.NamedArgs) {
	query := `INSERT INTO orders VALUES (@ID, @TrackNumber, @Entry, @Delivery, 
						@Locale, @InternalSignature, @CustomerID, @DeliveryService, @Shardkey,
						@SmID, @DateCreated, @OofShard)`

	args := pgx.NamedArgs{
		"ID": or.ID,
		"TrackNumber": or.TrackNumber,
		"Entry": or.Entry,
		"Delivery": or.Delivery.Id,
		"Locale": or.Locale,
		"InternalSignature": or.InternalSignature,
		"CustomerID": or.CustomerID,
		"DeliveryService": or.DeliveryService,
		"Shardkey": or.Shardkey,
		"SmID": or.SmID,
		"DateCreated": or.DateCreated,
		"OofShard": or.OofShard,
	}

	return query, args
}

func insertItemQuery(it entity.Item) (string, pgx.NamedArgs) {
	query := `INSERT INTO items VALUES (@ChrtID, @TrackNumber, @Price, @Rid, @Name, @Sale,
										@Size, @TotalPrice, @NmID, @Brand, @Status)`

	args := pgx.NamedArgs{
		"ChrtID": it.ChrtID,
		"TrackNumber": it.TrackNumber,
		"Price": it.Price,
		"Rid": it.Rid,
		"Name": it.Name,
		"Sale": it.Sale,
		"Size": it.Size,
		"TotalPrice": it.TotalPrice,
		"NmID": it.NmID,
		"Brand": it.Brand,
		"Status": it.Status,
	}

	return query, args
}

func (pg *OrderPostgres) GetById(orderId string) (entity.Order, error) {
	ctx := context.Background()
	var order entity.Order
	query := `SELECT * FROM orders o JOIN delivery d on o.delivery_id = d.id 
				JOIN payments p on o.payment_id = p.id 
				WHERE o.id = $1`
				//add select items
	err := pg.db.DB.QueryRow(ctx, query, orderId).Scan(&order)
	if err != nil {
		return order, err
	}
	
	return order, err
}

func (pg *OrderPostgres) GetAll() ([]entity.Order, error) {
	ctx := context.Background()

	query := `SELECT * FROM orders o`
	rows, _ := pg.db.DB.Query(ctx, query)
	orders, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Order])
	if err != nil {
		log.Printf("CollectRowsO: %v", err)
		return nil, err
	}	

	for i := range orders {
		queryDelivery := `SELECT * FROM delivery WHERE id = $1`
		row, _ := pg.db.DB.Query(ctx, queryDelivery, orders[i].DeliveryID)
		delivery, err := pgx.CollectOneRow(row, pgx.RowToStructByName[entity.Delivery])
		if err != nil {
			return nil, err
		}
		orders[i].Delivery = delivery

		queryPayment := `SELECT * FROM payments WHERE transaction = $1`
		row, _ = pg.db.DB.Query(ctx, queryPayment, orders[i].ID)
		payment, err := pgx.CollectOneRow(row, pgx.RowToStructByName[entity.Payment])
		if err != nil {
			return nil, err
		}
		orders[i].Payment= payment

		query := `SELECT * FROM items WHERE track_number = $1`
		rows, _ := pg.db.DB.Query(ctx, query, orders[i].TrackNumber)
		items, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Item])
		if err != nil {
			return nil, err
		}

		orders[i].Items = items
	}

	log.Printf("Scan: %v", orders[0])

	return orders, err
}
