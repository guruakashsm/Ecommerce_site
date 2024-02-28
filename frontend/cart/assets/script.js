var urlParams = new URLSearchParams(window.location.search);
var token = urlParams.get('token');
const productContainer = document.getElementById('product-container');
const totalCostValue = document.getElementById('total-cost-value');
const buyNowButton = document.getElementById('buy-now-button');

const itemsInCart = []; // Array to store items to be bought

// Function to fetch and display products
async function fetchAndDisplayProducts(token) {
    console.log("111")
    try {
        const data = {
            token: token
        };
        const response = await fetch('http://localhost:8080/products', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });// Replace with your backend API endpoint
        const products = await response.json();

        // Loop through the products and create product cards
        products.forEach(product => {
            const productCard = document.createElement('div');
            productCard.classList.add('product');

            let quantity = product.quantity; // Initial quantity from /products
            let total = product.price * quantity; // Calculate total based on initial quantity

            productCard.innerHTML = `
                <h2>${product.name}</h2>
                <!-- Add a bin icon for deleting the product -->
                <i class="delete-product-icon fa fa-trash"></i>
                <div class="quantity">
                    <button class="decrement">-</button>
                    <span>${quantity}</span>
                    <button class="increment">+</button>
                </div>
                <p>Price: $<span class="price">${product.price.toFixed(2)}</span></p>
                <p>Total: $<span class="total">${total.toFixed(2)}</span></p>
            `;

            const incrementButton = productCard.querySelector('.increment');
            const decrementButton = productCard.querySelector('.decrement');
            const quantitySpan = productCard.querySelector('.quantity span');
            const priceSpan = productCard.querySelector('.price');
            const totalSpan = productCard.querySelector('.total');

            // Add event listeners for the delete icons
            const deleteIcon = productCard.querySelector('.delete-product-icon');
            deleteIcon.addEventListener('click', () => {
                const productName = product.name;
                const productQuantity = product.quantity
                deleteProduct(productName, productQuantity);
            });

            incrementButton.addEventListener('click', () => {
                quantity++;
                quantitySpan.textContent = quantity;
                total = product.price * quantity;
                totalSpan.textContent = total.toFixed(2);
                // Call the updateCart function to send updated information
                updateCart(product.name, total, quantity);
                // Recalculate the total cost
                calculateTotalCost();
            });

            decrementButton.addEventListener('click', () => {
                if (quantity > 1) {
                    quantity--;
                    quantitySpan.textContent = quantity;
                    total = product.price * quantity;
                    totalSpan.textContent = total.toFixed(2);
                    // Call the updateCart function to send updated information
                    updateCart(product.name, total, quantity);
                    // Recalculate the total cost
                    calculateTotalCost();
                }
            });

            productContainer.appendChild(productCard);
        });
        // Calculate the initial total cost
        calculateTotalCost();
    } catch (error) {
        console.error('Error fetching products:', error);
    }
}

async function deleteProduct(productName, productQuantity) {
    try {
        const data = {
            token: token,
            name: productName,
            quantity: productQuantity // Include quantity in the request data
        };

        const response = await fetch('http://localhost:8080/deleteproduct', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        const result = await response.json();
        if (result === true) {
            calculateTotalCost();
            window.location.reload();
            // Perform any additional actions you need on success
        } else {
            alert("No product found");
        }
    } catch (error) {
        console.error('Error deleting product:', error);
        // Show an error message
        alert('Error deleting product');
    }
}



// Function to send updated product information to the backend
async function updateCart(productName, productPrice, productQuantity) {
    try {
        const data = {
            token: token,
            name: productName,
            price: productPrice,
            quantity: productQuantity
        };

        const response = await fetch('http://localhost:8080/updatecart', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        if (response.ok) {
            console.log('Cart updated successfully:', data);
        } else {
            console.error('Failed to update cart:', response.statusText);
        }
    } catch (error) {
        console.error('Error updating cart:', error);
    }
}

// Function to calculate and display the total cost
function calculateTotalCost() {
    const totalCost = Array.from(document.querySelectorAll('.total'))
        .map(span => parseFloat(span.textContent))
        .reduce((acc, currentValue) => acc + currentValue, 0);
    totalCostValue.textContent = totalCost.toFixed(2);
}

// Fetch and display products when the page loads
fetchAndDisplayProducts(token);

buyNowButton.addEventListener('click', async () => {
    const totalAmount = parseFloat(totalCostValue.textContent);

    if (!isNaN(totalAmount) && totalAmount > 0) {
        // Move the token variable definition here, within the click event listener
        var token = urlParams.get('token');

        // Redirect to the "/order" route with the token in the URL parameter
        window.location.href = `/order?token=${token}`;
    } else {
        // Handle the case where the total amount is not valid or zero
        alert('Total amount is not valid. Please add items to your cart.');
    }
});
