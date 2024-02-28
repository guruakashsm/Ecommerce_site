var urlParams = new URLSearchParams(window.location.search);
var token = urlParams.get('token');
var items = urlParams.get('item')
const productContainer = document.getElementById('product-container');

// Function to fetch and display inventory items
async function fetchAndDisplayInventory() {
    try {
        const response = await fetch('http://localhost:8080/inventorydata'); // Replace with your backend API endpoint
        const inventoryItems = await response.json();

        // Loop through the inventory items and create product cards
        inventoryItems.forEach(item => {
            const productCard = document.createElement('div');
            productCard.classList.add('product');

            // Check if the item has an image
            const hasImage = item.image && item.image.length > 0;

            productCard.innerHTML = `
                <div class="product-info">
                    <h2>Item Name: ${item.itemname}</h2>
                    <p>SellerName: ${item.sellername}<p>
                    <p>Category: ${item.itemcategory}</p>
                    <p>Price: $${item.price.toFixed(2)}</p>
                    <p>Quantity: ${item.quantity}</p>
                    ${item.sellerquantity < 10 ? `<p>Stock Left: ${item.sellerquantity}</p>` : ''}
                </div>
                ${hasImage ? `<img src="data:image/jpeg;base64,${item.image}" alt="Item Image" class="product-image">` : ''}
                <div class="quantity">
                    <button class="add-to-cart-button">Add to Cart</button>
                </div>

            `;

            productContainer.appendChild(productCard);

            // Add event listener to the "Add to Cart" button
            const addToCartButton = productCard.querySelector('.add-to-cart-button');
            addToCartButton.addEventListener('click', () => {
                // Create an object with the product name and price
                const cartItem = {
                    token: token,
                    name: item.itemname,
                    price: item.price,
                    quantity: item.sellerquantity
                };

                // Send the cart item to the backend
                addToCart(cartItem);
            });
        });
    } catch (error) {
        console.error('Error fetching inventory:', error);
    }
}

// Function to send the cart item to the backend
async function addToCart(cartItem) {
    try {
        const response = await fetch('http://localhost:8080/addtocart', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(cartItem)
        });

        if (response.ok) {
            console.log('Item added to cart:', cartItem);
            // Optionally, you can display a success message or update the UI here
        } else {
            console.error('Failed to add item to cart:', response.statusText);
            // Handle the error or display an error message
        }
    } catch (error) {
        console.error('Error adding item to cart:', error);
    }
}
var urlParams = new URLSearchParams(window.location.search);
var items = urlParams.get('item');

// Create a JavaScript object to represent the data
var data = {
    productName: items
};

// Send the data to the backend route "/searchitems" using a POST request
fetch('http://localhost:8080/search', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
})
    .then(response => response.json())
    .then(data => {
        // Handle the response from the server
        console.log('Response from server:', data);
    })
    .catch(error => {
        console.error('Error:', error);
    });
// Fetch and display inventory items when the page loads
fetchAndDisplayInventory();