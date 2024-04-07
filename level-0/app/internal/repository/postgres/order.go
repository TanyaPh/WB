package postgres

import (
	"api/internal/entity"
	"api/pkg/postgres"
	"context"

	"github.com/jackc/pgx/v5"
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
		return err
	}
	defer tx.Rollback(ctx)	// Rollback the transaction if not committed

	queryDelivery, args := insertDeliveryQuery(order.Delivery)
	if err = tx.QueryRow(ctx, queryDelivery, args).Scan(&order.Delivery.Id); err != nil {
		return err
	}

	queryPayment, args := insertPaymentQuery(order.Payment)
	if err = tx.QueryRow(ctx, queryPayment, args).Scan(&order.Payment.Id); err != nil {
		return err
	}

	// insert items


	queryOrder, args := insertOrderQuery(order)
	if err = tx.QueryRow(ctx, queryOrder, args).Scan(&order.ID); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func insertDeliveryQuery(d entity.Delivery) (string, pgx.NamedArgs) {
	query := "INSERT INTO delivery VALUES (@Name, @Phone, @Zip, @City, @Address, @Region, @Email) RETURNING id"
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
						@Amount, @PaymentDT, @Bank, @DeliveryCost, @GoodTotal, @CustomFree) 
					RETURNING id`
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
	query := `INSERT INTO payments VALUES (@ID, @TrackNumber, @Entry, @Delivery, @Payment, 
						@Locale, @InternalSignature, @CustomerID, @DeliveryService, @Shardkey
						@SmID, @DateCreated, @OofShard) 
					RETURNING id`
	args := pgx.NamedArgs{
		"ID": or.ID,
		"TrackNumber": or.TrackNumber,
		"Entry": or.Entry,
		"Delivery": or.Delivery.Id,
		"Payment": or.Payment.Id,
		// "Items": p.PaymentDT,
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
	var orders []entity.Order
	query := `SELECT * FROM orders o JOIN delivery d on o.delivery_id = d.id 
				JOIN payments p on o.order_uid = p.transaction`
				//JOIN items it on o.track_number = it.track_number 
				// add select items
	rows, err := pg.db.DB.Query(ctx, query)
	if err != nil {
		return orders, err
	}

	for rows.Next() {
		var order entity.Order
		
		if err := rows.Scan(&order); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}
	
	return orders, err
}