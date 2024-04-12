async function  fetchOrderInfo() {
    const orderId = document.getElementById('orderId').value;
    // console.log(orderId);

    const response = await fetch('http://localhost:8080/api/orders/' + orderId);
    if (response.ok) {
        const data = await response.json();
        document.getElementById('orderInfo').innerText = `ORDER ID: ${orderId}\n\n${JSON.stringify(data)}`
    } else {
        document.getElementById('orderInfo').innerText = `ORDER ID: ${orderId}\n\nError\nStatus: ${response.statusText}`
    }
  
}
