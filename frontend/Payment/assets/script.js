
async function fetchUserInfo() {
    try {
        const urlParams = new URLSearchParams(window.location.search);
        const token = urlParams.get('token'); // Extract the token from the URL

        if (token) {
            const data = {
                token: token
            };

            const response = await fetch('http://localhost:8080/getuser', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            }); // Send a POST request with the token in the request body

            if (response.ok) {
                const userData = await response.json();

                // Display user information in the HTML elements
                document.getElementById('firstname').textContent = userData.firstname;
                document.getElementById('lastname').textContent = userData.lastname;
                document.getElementById('phonenumber').textContent = userData.phonenumber;
                document.getElementById('houseno').textContent = userData.houseno;
                document.getElementById('streetname').textContent = userData.streetname;
                document.getElementById('city').textContent = userData.city;
                document.getElementById('pincode').textContent = userData.pincode;
            } else {
                console.error('Failed to fetch user information:', response.statusText);
            }
        } else {
            console.error('Token is missing in the URL.');
        }
    } catch (error) {
        console.error('Error fetching user information:', error);
    }
}
async function fetchTotalAmount() {
    try {
        const urlParams = new URLSearchParams(window.location.search);
        const token = urlParams.get('token'); // Extract the token from the URL

        if (token) {
            const data = {
                token: token
            };

            const response = await fetch('http://localhost:8080/totalamount', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            }); // Send a POST request with the token in the request body

            if (response.ok) {
                const userData = await response.json();
                document.getElementById('amount').textContent = userData.totalamount;
              
            } else {
                console.error('Failed to fetch user information:', response.statusText);
            }
        } else {
            console.error('Token is missing in the URL.');
        }
    } catch (error) {
        console.error('Error fetching user information:', error);
    }
}
fetchUserInfo();
fetchTotalAmount();