document.addEventListener("DOMContentLoaded", function () {
    const feedbackForm = document.getElementById("feedback-form");

    feedbackForm.addEventListener("submit", function (event) {
        event.preventDefault();

        const role = document.querySelector("input[name='role']:checked").value;
        const email = document.getElementById("email").value;
        const feedback = document.getElementById("feedback").value;

        const formData = {
            role: role, // Include the selected role in the form data
            email: email,
            feedback: feedback
        };


        // Send the formData to the "/feedback" backend route using an HTTP request
        fetch("http://localhost:8080/sitefeedback", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(formData)
        })
            .then(response => response.json())

            .then(data => {
                if (data === 1) {
                    // Redirect to /signin if the response is true
                    alert("Thankyou For Your WonderFull Feedback")
                    window.location.href = "/feedback";
                    document.querySelector("input[name='role']:checked").value = "";
                    document.getElementById("email").value = "";
                    document.getElementById("feedback").value = "";
                }
                else if (data === 0) {
                    // Handle other responses here, e.g., show an error message
                    alert("Error in Submitting");
                }
                else if (data === 2) {
                    alert("Email not Found")
                }
            })
            .catch(error => {
                console.error("Error:", error);
            });
    });
});