async function  fetchOrderInfo() {
    const orderId = document.getElementById('orderId').value;
    // console.log(orderId);
    document.getElementById('orderInfo').innerText = 'YOUR ORDER ID: ' + orderId;

    const response = await fetch('/api/order/' + orderId);
    if (response.ok) {
        const data = await response.json();
        document.getElementById('orderInfo').innerText = JSON.stringify(data);
    } else {
        document.getElementById('orderInfo').innerText = `Error\n Status: ${ response.statusText}`;
    }
  
}