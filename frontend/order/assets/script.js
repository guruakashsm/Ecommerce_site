var urlParams = new URLSearchParams(window.location.search);
var token = urlParams.get('token');
const ordersContainer = document.getElementById("ordersContainer");

// Function to display orders
function displayOrders(ordersData) {
    // Clear the previous orders if any
    ordersContainer.innerHTML = "";

    if (ordersData.length === 0) {
        // Handle the case when there are no orders
        const noOrdersMessage = document.createElement("p");
        noOrdersMessage.textContent = "No orders available.";
        ordersContainer.appendChild(noOrdersMessage);
    } else {
        // Iterate through the orders and create HTML elements
        ordersData.forEach(order => {
            const orderDiv = document.createElement("div");
            orderDiv.classList.add("order-item");

            const contentDiv = document.createElement("div"); // Content container

            const orderId = document.createElement("p");
            orderId.textContent = `Order ID: ${order._id}`;
            orderId.classList.add("order-info"); // Apply style to order ID

            const customerInfo = document.createElement("p");
            customerInfo.textContent = `Customer: ${order.address.firstname} ${order.address.lastname}, Phone: ${order.address.phonenumber}, City: ${order.address.city}, Pincode: ${order.address.pincode}`;
            customerInfo.classList.add("customer-info"); // Apply style to customer info

            const itemList = document.createElement("ul");
            order.itemstobuy.forEach(item => {
                const listItem = document.createElement("li");
                listItem.textContent = `${item.name}: ${item.quantity}`;
                listItem.classList.add("item-list"); // Apply style to item list
                itemList.appendChild(listItem);
            })

            const statusDiv = document.createElement("div");
            statusDiv.classList.add("status");
            statusDiv.innerHTML = '<div class="status-dot"></div> In Process';

            const deleteButton = document.createElement("button"); // Use a button element
            deleteButton.innerHTML = '<span class="icon">❎</span>Cancel Order'; // Insert an icon inside the button
            deleteButton.classList.add("delete-button");
            deleteButton.addEventListener("click", () => {
                // Prepare the data to send in the request body
                const data = {
                    _id: order._id,
                };

                // Send a DELETE request to the backend to delete the order
                fetch("http://localhost:8080/deleteorder", {
                    method: "DELETE",
                    headers: {
                        "Content-Type": "application.json",
                    },
                    body: JSON.stringify(data),
                })
                    .then(response => response.json())
                    .then(data => {
                        if (data.success) {
                            // Order was successfully deleted
                            alert("Order Canceled successfully.");
                            // Remove the deleted order from the display
                            orderDiv.remove();
                        } else {
                            // Handle the case where the deletion was not successful
                            alert("Error Concelling order. Please try again.");
                        }
                    })
                    .catch(error => {
                        // Handle errors, e.g., display an error message
                        alert("Error deleting order: " + error.message);
                    });
            });

            contentDiv.appendChild(orderId);
            contentDiv.appendChild(customerInfo);
            contentDiv.appendChild(itemList);

            orderDiv.appendChild(contentDiv);
            orderDiv.appendChild(statusDiv);
            orderDiv.appendChild(deleteButton);

            ordersContainer.appendChild(orderDiv);
        });
    }
}

// Fetch orders data and display it
const requestData = {
    token: token
};

// Send a POST request to the backend to fetch orders data
fetch("http://localhost:8080/customerorders", {
    method: "POST", // Use the POST method
    headers: {
        "Content-Type": "application/json"
    },
    body: JSON.stringify(requestData)
})
    .then(response => response.json())
    .then(data => {
        // Handle the orders data and display it
        displayOrders(data);
    })
    .catch(error => {
        console.error("Error fetching orders:", error);
    });