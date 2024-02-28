const loginForm = document.getElementById("login-form");
const dashboard = document.getElementById("dashboard");
const inventoryForm = document.getElementById("inventory-form");
const deleteProductForm = document.getElementById("delete-product-form");
const updateProductForm = document.getElementById("update-product-form");
const emailInput = document.getElementById("email");
const passwordInput = document.getElementById("password");
const loginButton = document.getElementById("login-button");
const createProductButton = document.getElementById("create-product-button");
const deleteProductButton = document.getElementById("delete-product-button");
const updateProductButton = document.getElementById("update-product-button");
const ordersButton = document.getElementById("orders-button");
const ordersContainer = document.getElementById("orders-container");

// Event listener for the login button
loginButton.addEventListener("click", () => {
    const email = emailInput.value;
    const password = passwordInput.value;

    // Create a structured object with email and password
    const loginData = {
        email: email,
        password: password
    };

    // Send a POST request to the backend route "/sellercheck" with the structured object
    fetch("http://localhost:8080/sellercheck", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(loginData)
    })
        .then(response => response.json())
        .then(data => {
            if (data.token) {
                // Redirect to the "/home" URL with the received token
                token = data.token;
                alert("Login Success");
                loginForm.style.display = "none";
                dashboard.style.display = "block";
            } else {
                // Handle the case where login was not successful
                alert("Login failed. Please check your credentials.");
            }
        })
        .catch(error => {
            console.error("Error:", error);
        });
});

// Event listener for the "Delete Product" button
deleteProductButton.addEventListener("click", () => {
    // Show the delete product form and hide other forms
    inventoryForm.style.display = "none";
    deleteProductForm.style.display = "block";
    updateProductForm.style.display = "none";
    ordersContainer.style.display = "none"
});

// Event listener for the "Create Product" button
createProductButton.addEventListener("click", () => {
    // Show the inventory form and hide other forms
    deleteProductForm.style.display = "none";
    inventoryForm.style.display = "block";
    updateProductForm.style.display = "none";
    ordersContainer.style.display = "none"
});

ordersButton.addEventListener("click", () => {
    deleteProductForm.style.display = "none";
    inventoryForm.style.display = "none";
    updateProductForm.style.display = "none";
    ordersContainer.style.display = "block"
})

updateProductButton.addEventListener("click", () => {
    // Hide other forms
    inventoryForm.style.display = "none";
    deleteProductForm.style.display = "none";
    // Show the update product form
    updateProductForm.style.display = "block";
    ordersContainer.style.display = "none"
});

// JavaScript to handle form submission for deleting a product
document.getElementById("deleteProductForm").addEventListener("submit", function (event) {
    event.preventDefault();
    const productNameToDelete = document.getElementById("productNameToDelete").value;

    // Create an object with the product name to delete
    const deleteProductData = {
        productname: productNameToDelete
    };

    // Send a POST request to the backend route "/deleteproduct" with the product name to delete
    fetch("http://localhost:8080/deleteproductbyseller", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(deleteProductData)
    })
        .then(response => response.json())
        .then(data => {
            if (data === 1) {
                // Handle success, e.g., show a success message
                alert("Product deleted successfully.");
                document.getElementById("productNameToDelete").value = "";
            } else if (data === 0) {
                // Handle the case where deletion was not successful, e.g., show an error message
                alert("Error deleting product. Product not found.");
            } else {
                // Handle other responses as needed
                alert("An error occurred while deleting the product.");
            }
        })
        .catch(error => {
            // Handle errors, e.g., display an error message
            alert("Error deleting product: " + error.message);
        });

    // Clear the input field
    document.getElementById("productNameToDelete").value = "";
});


ordersButton.addEventListener("click", () => {
    // Send a GET request to the backend to fetch orders data
    fetch("http://localhost:8080/orders")
        .then(response => response.json())
        .then(data => {
            // Handle the orders data and display it
            displayOrders(data);
        })
        .catch(error => {
            console.error("Error fetching orders:", error);
        });
});
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
            orderDiv.style.border = "1px solid #ccc"; // Add a border
            orderDiv.style.padding = "10px"; // Add some padding
            orderDiv.style.marginBottom = "10px"; // Add space between boxes
            orderDiv.style.display = "flex"; // Use flexbox to control the layout
            orderDiv.style.alignItems = "flex-start"; // Align items to the top

            const contentDiv = document.createElement("div"); // Content container
            contentDiv.style.flex = "1"; // Make it expand to fill available space

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
            });

            const deleteButton = document.createElement("button"); // Use a button element
            deleteButton.innerHTML = '<span class="icon">✅</span> Completed'; // Insert an icon inside the button
            deleteButton.classList.add("delete-button");
            deleteButton.style.fontSize = "14px"; // Reduce the button font size
            deleteButton.style.backgroundColor = "green"; // Change the button background color
            deleteButton.style.color = "white"; // Change the button text color
            deleteButton.style.border = "none"; // Remove button border
            deleteButton.style.padding = "5px 10px"; // Add padding to the button
            deleteButton.addEventListener("click", () => {
                // Prepare the data to send in the request body
                const data = {
                    _id: order._id, // Assuming _id is the identifier for the order
                };

                // Send a DELETE request to the backend to delete the order
                fetch("http://localhost:8080/deleteorder", {
                    method: "DELETE",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify(data),
                })
                    .then(response => response.json())
                    .then(data => {
                        if (data.success) {
                            // Order was successfully deleted
                            alert("Order Completed successfully.");
                            // Remove the deleted order from the display
                            orderDiv.remove();
                        } else {
                            // Handle the case where the deletion was not successful
                            alert("Error Completing order. Please try again.");
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
            orderDiv.appendChild(deleteButton);

            ordersContainer.appendChild(orderDiv);
        });
    }
}






// JavaScript to handle form submission for updating a product
document.getElementById("updateProductForm").addEventListener("submit", function (event) {
    event.preventDefault();
    const productNameToUpdate = document.getElementById("productNameToUpdate").value;
    const attributeToUpdate = document.getElementById("attributeToUpdate").value;
    let newValue = document.getElementById("newValue").value;

    // Depending on the attribute type, parse the new value accordingly
    if (attributeToUpdate === "price" || attributeToUpdate === "quantity") {
        newValue = parseInt(newValue);
    }

    // Create an object with the product name, attribute to update, and new value
    const updateProductData = {
        productname: productNameToUpdate,
        attribute: attributeToUpdate,
        newvalue: newValue
    };

    // Send a POST request to the backend route "/updatebyseller" with the update data
    fetch("http://localhost:8080/updateproductbyseller", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(updateProductData)
    })
        .then(response => response.json())
        .then(data => {
            if (data === 1) {
                // Handle success, e.g., show a success message
                alert("Product updated successfully.");
                document.getElementById("productNameToUpdate").value = "";
                document.getElementById("attributeToUpdate").value = "";
                document.getElementById("newValue").value = "";

            } else if (data === 0) {
                // Handle the case where updating was not successful, e.g., show an error message
                alert("Error updating product. Product not found or invalid attribute.");
            } else {
                // Handle other responses as needed
                alert("An error occurred while updating the product.");
            }
        })
        .catch(error => {
            // Handle errors, e.g., display an error message
            alert("Error updating product: " + error.message);
        });

    // Clear the input fields
    document.getElementById("productNameToUpdate").value = "";
    document.getElementById("newValue").value = "";
});




// JavaScript to handle form submission for creating a product (in inventory)
document.getElementById("inventoryForm").addEventListener("submit", function (event) {
    event.preventDefault();
    const imageFile = document.getElementById("image").files[0];
    const reader = new FileReader();

    reader.onload = function () {
        const base64Image = btoa(new Uint8Array(reader.result).reduce((data, byte) => data + String.fromCharCode(byte), ''));
        const formData = {
            sellerid: token,
            itemcategory: document.getElementById("itemCategory").value,
            itemname: document.getElementById("itemName").value,
            price: parseFloat(document.getElementById("price").value),
            quantity: document.getElementById("quantity").value,
            image: base64Image,
            sellerquantity: parseInt(document.getElementById("sellerquantity").value, 10),
        };

        // Send a POST request to your Go backend
        fetch("http://localhost:8080/inventory", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(formData),
        })
            .then(response => response.json())
            .then(data => {
                if (data.data === true) {
                    alert("Item Inserted in inventory");
                    document.getElementById("itemCategory").value = "";
                    document.getElementById("itemName").value = "";
                    document.getElementById("price").value = "";
                    document.getElementById("quantity").value = "";
                    document.getElementById("image").value = "";
                    document.getElementById("sellerquantity").value = "";
                } else {
                    alert("Item Name Already Exists");
                }
            })
            .catch(error => {
                // Handle errors, e.g., display an error message
                alert("Error creating customer profile: " + error.message);
            });
    };

    reader.readAsArrayBuffer(imageFile);
});
