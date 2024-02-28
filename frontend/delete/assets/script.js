const loginForm = document.getElementById("login-form");
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

function updateDeleteByLabel() {
    const selectedCollection = collectionSelect.value;
    const label = collectionOptions[selectedCollection] || "";

    // Update the label and placeholder text based on the selected collection
    deleteByLabel.textContent = `Delete by: ${label}`;
    idInput.placeholder = `Enter ${label}`;

    // Clear the input field
    idInput.value = "";
}

// Event listener for collection select
collectionSelect.addEventListener("change", updateDeleteByLabel);

loginButton.addEventListener("click", function () {
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    // Check username and password here (you can modify this condition)
    if (username === "admin" && password === "password") {
        loginForm.style.display = "none";
        deleteForm.style.display = "block";
        updateDeleteByLabel();
    } else {
        alert("Invalid username or password. Please try again.");
    }
});

document.getElementById("delete-form").addEventListener("submit", function (event) {
    event.preventDefault();

    const collection = document.getElementById("collection").value;
    const idValue = document.getElementById("id").value;

    const requestData = {
        collection: collection,
        idValue: idValue
    };

    fetch("http://localhost:8080/deletedata", {
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
            resultDiv.innerHTML = "<p>Delete successful.</p>";
        } else {
            resultDiv.innerHTML = "<p>Delete failed.</p>";
        }
    })
    .catch(error => {
        const resultDiv = document.getElementById("result");
        resultDiv.innerHTML = `<p>Error: ${error.message}</p>`;
    });
});