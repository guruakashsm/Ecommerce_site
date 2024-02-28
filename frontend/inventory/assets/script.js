document.getElementById("inventoryForm").addEventListener("submit", function (event) {
    event.preventDefault();
    const imageFile = document.getElementById("image").files[0];
    const reader = new FileReader();

    reader.onload = function () {
        const base64Image = btoa(new Uint8Array(reader.result).reduce((data, byte) => data + String.fromCharCode(byte), ''));
        const formData = {
            itemcategory: document.getElementById("itemCategory").value,
            itemname: document.getElementById("itemName").value,
            price: parseFloat(document.getElementById("price").value),
            quantity: document.getElementById("quantity").value,
            image: base64Image,
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
            if (data === 1) {
                // Redirect to /additems if the response is true
                alert("Item Inserted in inventory");
                window.location.href = "/additems";
            } 
            if (data === 0) {
                // Handle other responses here, e.g., show an error message
                alert("Error creating");
            }
            if (data === 2){
                alert("ItemName already exists");
            }

        })
        .catch(error => {
            // Handle errors, e.g., display an error message
            alert("Error creating customer profile: " + error.message);
        });
    };

    reader.readAsArrayBuffer(imageFile);
});