 // Function to hide the currently displayed result
 function hideResult() {
    const resultContainer = document.getElementById('result-container');
    resultContainer.style.display = 'none';
    const dataList = document.getElementById('data-list');
    dataList.innerHTML = '';
}

const createSellerButton = document.getElementById('create-seller-button');
const createSellerForm = document.getElementById('create-seller-form');
const feedbackForm = document.getElementById('feedback-form-container')
const feedbackbutton = document.getElementById('feedback-button')
const editbutton = document.getElementById('edit-button')

feedbackbutton.addEventListener('click', () => {
    updateFormElement.style.display = "none";
    feedbackForm.style.display = "block";
    deleteFormContainer.style.display = 'none';
    createSellerForm.style.display = 'none';
    hideResult(); // Hide previous result, if any

});

// Show the create seller form when the "Create Seller" button is clicked
createSellerButton.addEventListener('click', () => {
    updateFormElement.style.display = "none";
    createSellerForm.style.display = 'block';
    hideResult(); // Hide previous result, if any
    deleteFormContainer.style.display = 'none';
    feedbackForm.style.display = 'none';
});

const loginForm = document.getElementById('login-form');
const optionsDiv = document.querySelector('.options');
const getAllDataButton = document.getElementById('get-all-data');
const getInventoryButton = document.getElementById('get-inventory');
const getSellerButton = document.getElementById('get-seller');

loginForm.addEventListener('submit', (e) => {
    e.preventDefault();
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    if (username === 'admin' && password === 'password') {
        alert('Login successful!');
        hideResult(); // Hide previous result, if any
        optionsDiv.style.display = 'block';
        loginForm.style.display = 'none'; // Hide login form
    } else {
        alert('Invalid username or password. Please try again.');
    }

});



const updateFormElement = document.getElementById("update-form-admin");
const collectionSelectElement = document.getElementById("updatecollection");
const fieldSelectElement = document.getElementById("field");

const collectionselectOptions = {
    customer: ["name", "email", "phonenumber", "age", "password", "firstname", "lastname", "houseno", "streetname", "city", "pincode"],
    inventory: ["itemcategory", "itemname", "price", "quantity"],
    seller: ["sellername", "selleremail", "password", "phoneno", "address"],
};

function populateFieldOptions() {
    const selectedCollection = collectionSelectElement.value;
    const options = collectionselectOptions[selectedCollection] || [];

    // Clear existing options
    fieldSelectElement.innerHTML = "";

    // Add new options
    options.forEach(option => {
        const optionElement = document.createElement("option");
        optionElement.value = option;
        optionElement.textContent = option;
        fieldSelectElement.appendChild(optionElement);
    });
}

// Event listener for collection select
collectionSelectElement.addEventListener("change", populateFieldOptions);

editbutton.addEventListener('click', () => {
    updateFormElement.style.display = "block";
    feedbackForm.style.display = "none";
    deleteFormContainer.style.display = 'none';
    createSellerForm.style.display = 'none';
    hideResult(); // Hide previous result, if any
});


document.getElementById("update-form").addEventListener("submit", function (event) {
    event.preventDefault();

    const updatecollection = document.getElementById("updatecollection").value;
    const idname = document.getElementById("idname").value;
    const field = document.getElementById("field").value;
    const newvalue = document.getElementById("newvalue").value;

    const requestData = {
        collection: updatecollection,
        email: idname,
        field: field,
        newvalue: newvalue
    };

    fetch("https://localhost:8080/update", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(requestData)
    })
        .then(response => response.json())
        .then(data => {
            const resultDiv = document.getElementById("result");
            if (data) {
                resultDiv.innerHTML = "<p>Update successful.</p>";
                document.getElementById("update-form").reset();
            } else {
                resultDiv.innerHTML = "<p>Update failed.</p>";
                document.getElementById("update-form").reset();
            }
        })
        .catch(error => {
            const resultDiv = document.getElementById("result");
            resultDiv.innerHTML = `<p>Error: ${error.message}</p>`;
        });
});




getAllDataButton.addEventListener('click', () => {
    hideResult(); // Hide previous result, if any
    createSellerForm.style.display = 'none'; // Close create seller form
    feedbackForm.style.display = 'none';
    deleteFormContainer.style.display = 'none';
    updateFormElement.style.display = "none";
    fetch('http://localhost:8080/getallcustomerdata')
        .then(response => response.json())
        .then(data => {
            const resultContainer = document.getElementById('result-container');
            resultContainer.style.display = 'block';
            const dataList = document.getElementById('data-list');
            dataList.innerHTML = '';

            data.forEach(customer => {
                const customerBox = document.createElement('div');
                customerBox.classList.add('customer-box');
                customerBox.innerHTML = `
                    <h3>Name: ${customer.name}</h3>
                    <p>Email: ${customer.email}</p>
                    <p>Phone No: ${customer.phonenumber}</p>
                    <p>Age: ${customer.age}</p>
                    <p>Password: ${customer.password}</p>
                    <p>First Name: ${customer.firstname}</p>
                    <p>Last Name: ${customer.lastname}</p>
                    <p>House No: ${customer.houseno}</p>
                    <p>Street Name: ${customer.streetname}</p>
                    <p>City: ${customer.city}</p>
                    <p>Pincode: ${customer.pincode}</p>
                `;
                dataList.appendChild(customerBox);
            });
        })
        .catch(error => {
            console.error('Error fetching data:', error);
        });
});

// Handle the form submission for creating a seller
const sellerForm = document.getElementById('seller-form');
sellerForm.addEventListener('submit', (e) => {
    e.preventDefault();
    hideResult(); // Hide previous result, if any
    createSellerForm.style.display = 'none'; // Close create seller form
    deleteFormContainer.style.display = 'none';

    // Collect seller data from the form
    const sellerData = {
        sellername: document.getElementById('seller-name').value,
        selleremail: document.getElementById('seller-email').value,
        password: document.getElementById('seller-password').value,
        confirmpassword: document.getElementById('seller-confirm-password').value,
        phoneno: parseInt(document.getElementById('seller-phone').value),
        address: document.getElementById('seller-address').value
    };

    // Send the seller data as JSON in the request body
    fetch('http://localhost:8080/createseller', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(sellerData)
    })
        .then(response => response.json())
        .then(data => {
            // Handle the response if needed
            alert('Seller created successfully.');
            // Optionally, clear the form or perform other actions
            sellerForm.reset();
        })
        .catch(error => {
            console.error('Error creating seller:', error);
        });
});
getSellerButton.addEventListener('click', () => {
    hideResult(); // Hide previous result, if any
    createSellerForm.style.display = 'none'; // Close create seller form
    deleteFormContainer.style.display = 'none';
    feedbackForm.style.display = 'none';
    updateFormElement.style.display = "none";
    fetch('http://localhost:8080/getallsellerdata')
        .then(response => response.json())
        .then(data => {
            console.log(data)
            const resultContainer = document.getElementById('result-container');
            resultContainer.style.display = 'block';
            const dataList = document.getElementById('data-list');
            dataList.innerHTML = '';

            data.forEach(seller => {
                const sellerBox = document.createElement('div');
                sellerBox.classList.add('seller-box');
                sellerBox.innerHTML = `
            <h3>Seller Details</h3>
            <p><strong>Seller Id:</strong> ${seller.sellerid}</p>
            <p><strong>Seller Name:</strong> ${seller.sellername}</p>
            <p><strong>Seller Email:</strong> ${seller.selleremail}</p>
            <p><strong>Password:</strong> ${seller.password}</p>
            <p><strong>Phone No:</strong> ${seller.phoneno}</p>
            <p><strong>Address:</strong> ${seller.address}</p>
        `;
                dataList.appendChild(sellerBox);
            });
        })
        .catch(error => {
            console.error('Error fetching seller data:', error);
        });
});

// Event listener for "Get Inventory Data" button
getInventoryButton.addEventListener('click', () => {

    hideResult(); // Hide previous result, if any
    createSellerForm.style.display = 'none';
    deleteFormContainer.style.display = 'none'; // Close create seller form
    feedbackForm.style.display = 'none';
    updateFormElement.style.display = "none";

    fetch('http://localhost:8080/getallinventorydata')
        .then(response => response.json())
        .then(data => {
            const resultContainer = document.getElementById('result-container');
            resultContainer.style.display = 'block';
            const dataList = document.getElementById('data-list');
            dataList.innerHTML = '';

            data.forEach(inventoryItem => {
                const inventoryBox = document.createElement('div');
                inventoryBox.classList.add('inventory-box');
                inventoryBox.innerHTML = `
                    <h3>Item Name: ${inventoryItem.itemname}</h3>
                    <p>Category: ${inventoryItem.itemcategory}</p>
                    <p>Price: $${inventoryItem.price.toFixed(2)}</p>
                    <p>Quantity: ${inventoryItem.quantity}</p>
                `;
                dataList.appendChild(inventoryBox);
            });
        })
        .catch(error => {
            console.error('Error fetching data:', error);
        });
});



const deleteForm = document.getElementById("delete-form");
const loginButton = document.getElementById("login-button");
const collectionSelect = document.getElementById("collection");
const deleteByLabel = document.getElementById("delete-by-label");
const idInput = document.getElementById("id");




const collectionOptions = {
    customer: "Email",
    inventory: "Item Name",
    seller: "Email"
};

const deleteFormContainer = document.getElementById("delete-form-container"); // Container for delete form
const deleteButton = document.getElementById("delete-button"); // Corrected button id
const resultDiv = document.getElementById("result-container"); // Result container

// Show the delete form when the "Delete" button is clicked
deleteButton.addEventListener("click", () => {
    updateFormElement.style.display = "none";
    deleteFormContainer.style.display = 'block';
    createSellerForm.style.display = 'none';
    feedbackForm.style.display = 'none';

    hideResult(); // Hide previous result, if any
});

// Event listener for delete form submission
document.getElementById("delete-form").addEventListener("submit", function (event) {
    event.preventDefault();

    const collection = document.getElementById("collection").value;
    const idValue = document.getElementById("id").value;

    const requestData = {
        collection: collection,
        idValue: idValue
    };

    // Send a DELETE request to your server to delete the data
    fetch("http://localhost:8080/deletedata", {
        method: "POST", // Use DELETE method to delete data
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(requestData)
    })
        .then(response => response.json())
        .then(data => {
            const resultDiv = document.getElementById("result");
            if (data === true) {
                alert("Deleted Sucessfull")
                document.getElementById("id").value = "";
            } else {
                alert("Error in Deleting")
            }
        })
        .catch(error => {
            const resultDiv = document.getElementById("result-container");
            resultDiv.innerHTML = `<p>Error: ${error.message}</p>`;
        });
});



document.addEventListener("DOMContentLoaded", function () {
    const sellerButton = document.getElementById("sellerButton");
    const customerButton = document.getElementById("customerButton");
    const feedbackContainer = document.getElementById("feedbackContainer");

    sellerButton.addEventListener("click", function () {
        fetch("http://localhost:8080/sellerfeedback")
            .then(response => response.json())
            .then(data => {
                displayFeedback(data, "Seller");
            })
            .catch(error => {
                console.error("Error:", error);
            });
    });

    customerButton.addEventListener("click", function () {
        fetch("http://localhost:8080/customerfeedback")
            .then(response => response.json())
            .then(data => {
                displayFeedback(data, "Customer");
            })
            .catch(error => {
                console.error("Error:", error);
            });
    });

    // Function to add a delete icon to each feedback item
    function addDeleteIcon(feedbackBox, email, feedback) {
        const deleteIcon = document.createElement("span");
        deleteIcon.classList.add("delete-icon");
        deleteIcon.innerHTML = "&#10006;"; // X icon
        deleteIcon.addEventListener("click", function () {
            // Send email and feedback to the "/deletefeedback" route
            fetch("http://localhost:8080/deletefeedback", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ email, feedback })
            })
                .then(response => response.json())
                .then(data => {
                    if (data === 1) {
                        // Feedback deleted successfully, remove the feedback item from the UI
                        feedbackBox.remove();
                    } else {
                        alert("Error deleting feedback");
                    }
                })
                .catch(error => {
                    console.error("Error:", error);
                });
        });
        feedbackBox.appendChild(deleteIcon);
    }

    // Function to display feedback with delete icons
    function displayFeedback(feedbackData, role) {
        feedbackContainer.innerHTML = "";
        feedbackData.forEach(item => {
            const feedbackBox = document.createElement("div");
            feedbackBox.classList.add("feedback-box");
            feedbackBox.innerHTML = `

                <p><strong>Email:</strong> ${item.email}</p>
                <p><strong>Feedback:</strong> ${item.feedback}</p>
            `;
            addDeleteIcon(feedbackBox, item.email, item.feedback);
            feedbackContainer.appendChild(feedbackBox);
        });
    }
});