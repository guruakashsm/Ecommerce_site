const sellerForm = document.getElementById('login-button');
sellerForm.addEventListener('click', (e) => {
  console.log("Clicker seller")
  e.preventDefault();

  const sellerData = {
    email: document.getElementById('seller-email').value,
    password: document.getElementById('seller-password').value,
  };

  if (sellerData.selleremail == "" || sellerData.password == "") {
    showToast("Please enter all fields before Submit", "Info", 1);
    return;
  }
  if (sellerData.password.length < 6) {
    showToast("Incorrect Password", "Danger", 0);
    return;
  }

  // Send the seller data as JSON in the request body
  fetch('http://localhost:8080/sellercheck', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(sellerData)
  })
    .then(response => response.json())
    .then(data => {
      if (data.message) {
        showToast(data.message, "Danger", 0)

      } else if (data.error) {
        showToast(data.error, 'Error', 0);
      } else if (data.token) {
        showToast("Login Successfull", "Success", 3)
        const sellerData = {
          'token': data.token,
          'username': document.getElementById('seller-email').value
        }
        const jsonString = JSON.stringify(sellerData);
        localStorage.setItem('sellerdata', `${jsonString}`);
        setTimeout(()=>{
          window.location.href = "/seller/dashboard"
        },2000)
      }

    })
    .catch(error => {
      showToast(error, 'Error', 0);
    });


});

function showToast(str, war, no) {
  const toastContainer = document.querySelector('.toast-container');
  const title = document.querySelector('.js-toast-title');
  const content = document.querySelector('.js-toast-content');
  const image = document.querySelector('.js-toast-img');

  // Reset classes, width, and height
  toastContainer.className = 'toast-container';
  toastContainer.style.width = 'auto';
  toastContainer.style.height = 'auto';

  if (no == 0) {
    image.src = './images/danger.webp';
    toastContainer.classList.add('danger-color');
  } else if (no == 1) {
    image.src = './images/info.svg';
    toastContainer.classList.add('info-color');
  } else if (no == 2) {
    image.src = './images/warning.jpg';
    toastContainer.classList.add('warning-color');
  } else if (no == 3) {
    image.src = './images/success.png';
    toastContainer.classList.add('success-color');
  }
  title.innerHTML = `${war}`;
  content.innerHTML = `${str}`;

  // Calculate and set the container width and height

  const containerWidth = title.length + content.length + 500; // Add some padding

  toastContainer.style.width = `${containerWidth}px`;


  // Add transition effect
  toastContainer.style.transition = 'all 0.5s ease-in-out';

  toastContainer.style.display = 'block';
  setTimeout(() => {
    toastContainer.style.opacity = 1;
  }, 1);

  // Hide the toast container after 5 seconds
  setTimeout(() => {
    toastContainer.style.opacity = 0;
    setTimeout(() => {
      toastContainer.style.display = 'none';
    }, transitionDuration * 1000);
  }, 3000);
}