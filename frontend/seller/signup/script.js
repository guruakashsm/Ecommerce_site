const sellerForm = document.getElementById('seller-form');
sellerForm.addEventListener('submit', (e) => {
    e.preventDefault();
    console.log("in Seller")
    const imageFile = document.getElementById("seller-image").files[0];
    const imageinput = document.getElementById("seller-image");
    if (imageinput.files.length === 0) {
        showToast("Please select an image", "Info", 1);
        return;
    }
    const reader = new FileReader();
    
    reader.onload = function () {
        const base64Image = btoa(new Uint8Array(reader.result).reduce((data, byte) => data + String.fromCharCode(byte), ''));// Extracting base64 data from data URL
        const sellerData = {
            sellername: document.getElementById('seller-name').value,
            selleremail: document.getElementById('seller-email').value,
            password: document.getElementById('seller-password').value,
            confirmpassword: document.getElementById('seller-confirm-password').value,
            phoneno: parseInt(document.getElementById('seller-phone').value),
            address: document.getElementById('seller-address').value,
            image: base64Image,
        };

        if (sellerData.sellername == "" || sellerData.selleremail == "" || sellerData.password == "" || sellerData.confirmpassword == "" || sellerData.phoneno == "" || sellerData.address == "") {
            showToast("Please enter all fields before Submit", "Info", 1);
            return;
        }
        if (sellerData.password != sellerData.confirmpassword) {
            showToast("Password Mismatch", "Danger", 0);
            return;
        }

        // Send the seller data as JSON in the request body
        fetch('http://localhost:8080/registerseller', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(sellerData)
        })
            .then(response => response.json())
            .then(data => {
              if(data.message ){
                showToast(data.message, 'Success', 3);
                if(data.message == "Verify Your Email"){
                  DisplayVerifyEmail()
                }else if(data.message == "Email already Exists and Verified"){
                      window.location.href = "/seller/signin"
                }

              }else if(data.error){
                showToast(data.error, 'Success', 3);
              }
               
            })
            .catch(error => {
                showToast(error, 'Error', 0);
            });
    };

    // Read the image file as a data URL
    reader.readAsArrayBuffer(imageFile);
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

  function DisplayVerifyEmail(){
    document.getElementById("message-form").style.display ='none';
    document.getElementById("otp-form").style.display = 'block';
    sellerForm.style.display= 'none'
  }


  document.getElementById("otp-verify").addEventListener("click",(event)=>{
    event.preventDefault();
    let otp = document.getElementById("otp").value
    if (otp === "" || otp.length != 6) {
        showToast("Please enter valid OTP", "Danger", 0);
        return false;
    }
    const formData = {
        email: document.getElementById("seller-email").value,
        verification: otp,
    };
    fetch("http://localhost:8080/verifyselleremail", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
    })
        .then(response => response.json())
        .then(data => {
            if (data.message == "Wrong OTP") {
                showToast(`${data.message}`, "Danger", 0);
            } else if(data.message == "Email Verification Successful") {
                showToast(`${data.message}`, "Success", 3);
                setTimeout(() => {
                    document.getElementById("seller-form").reset()
                    document.getElementById("otp-form").style.display = 'none';
                    document.getElementById("seller-form").style.display = 'none';
                    document.getElementById("message-form").style.display ='block';
                   
                }, 2000);
            }else if(data.error){
              showToast(`${data.error}`, "Danger", 0);
            }
        })
        .catch(error => {
            
            showToast(`Error: ${error.message}`, "Danger", 0);
        });
    return false; 
})
  
