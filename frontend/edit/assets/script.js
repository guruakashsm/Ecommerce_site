const loginForm = document.getElementById("login-form");
        const updateForm = document.getElementById("update-form");
        const loginButton = document.getElementById("login-button");
        const collectionSelect = document.getElementById("collection");
        const fieldSelect = document.getElementById("field");

        const collectionOptions = {
            customer: ["name", "email", "phonenumber","age","password","firstname","lastname","houseno","streetname","city","pincode"],
            inventory: ["itemcategory", "itemname", "price","quantity"],
            seller: ["sellername", "selleremail", "password","phoneno","address"],
        };

        function populateFieldOptions() {
            const selectedCollection = collectionSelect.value;
            const options = collectionOptions[selectedCollection] || [];

            // Clear existing options
            fieldSelect.innerHTML = "";

            // Add new options
            options.forEach(option => {
                const optionElement = document.createElement("option");
                optionElement.value = option;
                optionElement.textContent = option;
                fieldSelect.appendChild(optionElement);
            });
        }

        // Event listener for collection select
        collectionSelect.addEventListener("change", populateFieldOptions);

        loginButton.addEventListener("click", function () {
            const username = document.getElementById("username").value;
            const password = document.getElementById("password").value;
            
            // Check username and password here (you can modify this condition)
            if (username === "admin" && password === "password") {
                loginForm.style.display = "none";
                updateForm.style.display = "block";
                populateFieldOptions();
            } else {
                alert("Invalid username or password. Please try again.");
            }
        });
        
        document.getElementById("update-form").addEventListener("submit", function (event) {
            event.preventDefault();

            const collection = document.getElementById("collection").value;
            const idname = document.getElementById("idname").value;
            const field = document.getElementById("field").value;
            const newvalue = document.getElementById("newvalue").value;

            const requestData = {
                collection: collection,
                email: idname,
                field: field,
                newvalue: newvalue
            };
            
            fetch("http://localhost:8080/update", {
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