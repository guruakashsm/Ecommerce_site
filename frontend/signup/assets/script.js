document.getElementById("customerForm").addEventListener("submit", function (event) {
    event.preventDefault();

    // Create a JSON object from the form data
    const formData = {
        name: document.getElementById("name").value,
        email: document.getElementById("email").value,
        age: parseInt(document.getElementById("age").value),
        phonenumber: parseInt(document.getElementById("phonenumber").value),
        password: document.getElementById("password").value,
        confirmpassword: document.getElementById("confirmpassword").value,
        firstname: document.getElementById("firstname").value,
        lastname: document.getElementById("lastname").value,
        houseno: document.getElementById("houseno").value,
        streetname: document.getElementById("streetname").value,
        city: document.getElementById("city").value,
        pincode: parseInt(document.getElementById("pincode").value),
    };

    // Send a POST request to your Go backend
    fetch("http://localhost:8080/create", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
    })
    .then(response => response.json())
    
    .then(data => {
        if (data === 1) {
            // Redirect to /signin if the response is true
            alert("Signup Sucessfull")
            window.location.href = "/signin";
        } 
        else if (data === 0) {
            // Handle other responses here, e.g., show an error message
            alert("Error creating customer profile.");
        }
       else if (data === 2){
            alert("Email already exists")
        }
        else if (data === 3){
            alert("Password and Confirm Password defer")
        }
        else if (data === 4){
            alert("User name should only contain letters")
        }
        else if (data === 5){
            alert("Invalid Phonenumber")
        }
        else if (data === 6){
            alert("Invalid Pincode")
        }
    })
    .catch(error => {
        // Handle errors, e.g., display an error message
        alert("Error creating customer profile: " + error.message);
    });
});