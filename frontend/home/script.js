function HideHomePage() {
  document.querySelector('.blog').style.display = 'none'
  document.querySelector('.testimonial').style.display = 'none'
  document.querySelector('.product-container').style.display = 'none'
  document.querySelector('.category').style.display = 'none'
  document.querySelector('.banner').style.display = 'none'
  document.querySelector('.desktop-navigation-menu').style.display = 'none'
  document.querySelector('.cta-container').style.display = 'none'
  document.querySelector('.service').style.display = 'none'
}


function HomePage() {
  document.querySelector('.blog').style.display = 'block'
  document.querySelector('.testimonial').style.display = 'block'
  document.querySelector('.product-container').style.display = 'block'
  document.querySelector('.category').style.display = 'block'
  document.querySelector('.banner').style.display = 'block'
  document.querySelector('.desktop-navigation-menu').style.display = 'block'
  document.querySelector('.cta-container').style.display = 'block'
  document.querySelector('.service').style.display = 'block'
  document.getElementById('js-display-items').style.display = 'none'
  document.querySelector('.checkout-container').style.display = 'none'
  document.getElementById("single-order-container").style.display = 'none'
  document.getElementById("order-container").style.display = 'none'
  document.getElementById('checkout-container').style.display = 'none'

}




document.getElementById("searchBtn").addEventListener("click", function () {

  const productName = document.getElementById("searchField").value;
  const productNameInUpperCase = productName.toUpperCase();

  fetch("http://localhost:8080/search", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ productName: productNameInUpperCase }),
  })
    .then(response => response.json())
    .then(data => {

      let html = ""
      if (data.data) {
        data.data.forEach((item) => {
          html += ` 
                <div class="col-xs-12 col-md-6 bootstrap snippets bootdeys">

                <div class="product-content product-wrap clearfix">
                <div class="row">
                <div class="col-md-5 col-sm-12 col-xs-12">
                <div class="product-image">
                <img onclick="DisplayData('${item.itemname}')"src="data:image/jpeg;base64,${item.image}" alt="194x228" class="img-responsive">
                <span class="tag2 hot">
                HOT
                </span>
                </div>
                </div>
                <div class="col-md-7 col-sm-12 col-xs-12">
                <div class="product-deatil">
                <h5 class="name">
                <a href="#">
                ${item.itemname}  <span>${item.itemcategory}</span>
                </a>
                </h5>
                <p class="price-container">
                <span>₹${item.price.toFixed(2)}</span>
                </p>
                <span class="tag1"></span>
                </div>
                <div class="description">
                <p>${item.shortdis}</p>
                </div>
                <div class="product-info smart-form">
                <div class="row">
                <div class="col-md-6 col-sm-6 col-xs-6">
                <a onclick="AddtoCart('${item.itemname}')" class="btn btn-success" style="color:white;">Add to cart</a>
                </div>
                <div class="col-md-6 col-sm-6 col-xs-6">
                <div class="rating">
                <label for="stars-rating-5"><i class="fa fa-star"></i></label>
                <label for="stars-rating-4"><i class="fa fa-star"></i></label>
                <label for="stars-rating-3"><i class="fa fa-star text-primary"></i></label>
                <label for="stars-rating-2"><i class="fa fa-star text-primary"></i></label>
                <label for="stars-rating-1"><i class="fa fa-star text-primary"></i></label>
                </div>
                </div>
                </div>
                </div>
                </div>
                </div>
                </div>
                
                </div>
                 `
        })
      }
      if (html == "") {
        html = `<img style="margin-left:15%;"src="./assets/images/noresult.gif" alt="No Results Found">`
      }
      document.querySelector('.checkout-container').style.display = 'none'
      document.getElementById("single-order-container").style.display = 'none'
      document.getElementById("order-container").style.display = 'none'
      document.getElementById('checkout-container').style.display = 'none'
      document.getElementById("js-display-items").innerHTML = html;
      document.getElementById("js-display-items").style.display = 'block';
      HideHomePage()

    })
    .catch(error => {
      showToast(error, "Error", 0);
    });
});

function CheckNil() {
  const value = document.getElementById("searchField").value
  if (value.trim() == "") {
    HomePage()
  }
}

function AddtoCart(productName) {
  const storedData = localStorage.getItem("userdata");
  const retrievedUserData = JSON.parse(storedData);
  fetch("http://localhost:8080/addtocart", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ productName, token: retrievedUserData.token }),
  })
    .then(response => response.json())
    .then(data => {
      if (data.error) {
        showToast(data.error, "Danger", 0)
      } else if (data.message) {
        showToast(data.message, "Success", 3)
      }

    })
    .catch(error => {
      showToast(error, "Error", 0);
    });
}

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

function DisplayData(productName) {
  fetch("http://localhost:8080/getinventorydata", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ productName, }),
  })
    .then(response => response.json())
    .then(data => {

      let html = ""
      if (data.message) {
        html += `
                <div class="container">
                <img src="http://upload.wikimedia.org/wikipedia/commons/thumb/b/b1/Back_Arrow.svg/2048px-Back_Arrow.svg.png" height="35px" onclick="BackButton()" style="margin-left:-50px;cursor:pointer">
                  <!-- Left Column / Headphones Image -->
                  <div class="left-column">
                    <img data-image="red" class="active" src="data:image/jpeg;base64,${data.message.image}" alt="" height="400px" style="margin-right:100px">
                  </div>
              
                  <!-- Right Column -->
                  <div class="right-column">
              
                    <!-- Product Description -->
                    <div class="product-description">
                      <span>${data.message.itemcategory}</span>
                      <h1>${data.message.itemname}</h1>
                      <p>${data.message.longdis}</p>
                    </div>
              
                    <!-- Product Configuration -->
                    <div class="product-configuration">
              
                      <!-- Cable Configuration -->
                      <div class="cable-config">
                        <span>Features</span>
              
                        <div class="cable-choose">
                          
                        </div>
              
                        <a>${data.message.shortdis}</a>
                      </div>
                    </div>
              
                    <!-- Product Pricing -->
                    <div class="product-price">
                      <span>₹${data.message.price}</span>
                      <a onclick="AddtoCart('${data.message.itemname}')" class="cart-btn" style="color:white;">Add to cart</a>
                    </div>
                  </div>
              
                </div>
              `;

        function generateFeatureButtons(features) {
          // Map each feature to a button element
          const buttons = features.map(feature => `<button>${feature}</button>`).join('');
          return buttons;
        }


      }
      if (html == "") {
        html = `<img style="margin-left:15%;"src="./assets/images/noresult.gif" alt="No Results Found">`
      }
      document.querySelector('.checkout-container').style.display = 'none'
      document.getElementById("single-order-container").style.display = 'none'
      document.getElementById("order-container").style.display = 'none'
      document.getElementById('checkout-container').style.display = 'none'
      document.getElementById("js-display-items").innerHTML = html;
      document.getElementById("js-display-items").style.display = 'block';
      HideHomePage()

    })
    .catch(error => {
      showToast(error, "Error", 0);
    });
}

function BackButton() {
  document.getElementById("searchBtn").click()
}

function DisplayCart() {
  const storedData = localStorage.getItem("userdata");
  const retrievedUserData = JSON.parse(storedData);
  fetch("http://localhost:8080/products", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ token: retrievedUserData.token }),
  })
    .then(response => response.json())
    .then(data => {

      let html = ""
      let price = 0
      if (data.message) {
        html += `<div class="row">
        <div class="col-xl-8">`
        data.message.forEach((item) => {
          price += item.totalprice
          html += ` 
                <div class="card border shadow-none">
                <div class="card-body">
                  <div class="d-flex align-items-start border-bottom pb-3">
                    <div class="me-4">
                      <img src="data:image/jpeg;base64,${item.image}" alt class="avatar-lg rounded">
                    </div>
                    <div class="flex-grow-1 align-self-center overflow-hidden">
                      <div>
                        <h5 class="text-truncate font-size-18"><a href="#" class="text-dark">${item.productname.toUpperCase()}</a></h5>
      
                        <p class="mb-0 mt-1">Category : <span class="fw-medium">${item.itemcategory.toUpperCase()}</span></p>
                        <p class="text-muted mb-0" style="font-size: small; margin-top: 5px;">
                          Sold By: ${item.sellername}
                        </p>
                      </div>
                    </div>
                    <div class="flex-shrink-0 ms-2">
                      <ul class="list-inline mb-0 font-size-16">
                        <li class="list-inline-item">
                          <a href="#" class="text-muted px-1" onclick="DeleteProduct('${item.productname}','${item.quantity}')">
                            <i class="mdi mdi-trash-can-outline"></i>
                          </a>
                        </li>
                      </ul>
                    </div>
                  </div>
                  <div>
                    <div class="row">
                      <div class="col-md-4">
                        <div class="mt-3">
                          <p class="text-muted mb-2">Price</p>
                          <h5 class="mb-0 mt-2">₹${item.price}</h5>
                        </div>
                      </div>
                      <div class="col-md-5">
                        <div class="mt-3">
                          <p class="text-muted mb-2">Quantity</p>
                          <div class="d-inline-flex">
                            <div class="quantity">
                                <a class="btn btn-success decrement" style="display: inline;color:white;" onclick="DecrementButton('${item.productname}','${item.quantity}')" >-</a>
                              <span>${item.quantity}</span>
                                <a  class="btn btn-success increment" style="display: inline;color:white;" onclick="IncrementButton('${item.productname}','${item.quantity}')"> + </a>
                            </div>
                          </div>
                        </div>
                      </div>
                      <div class="col-md-3">
                        <div class="mt-3">
                          <p class="text-muted mb-2">Total</p>
                          <h5>₹${item.totalprice}</h5>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
      
                 `


        })
      }
      if (html == "") {
        html = `
              <div class="row my-4">
              <div class="col-sm-6">
                 <a onclick="HomePage()" class="btn btn-link text-muted">
                 <i class="mdi mdi-arrow-left me-1"></i> Continue Shopping </a>
              </div>
              </div>
              <img style="margin-left:20%; height:500px"src="./images/empty.png" alt="No Results Found">
              <h9 style="margin-left:40%;">Oops!! Your Cart is Empty.</h9>          
            `
        document.querySelector('.checkout-container').style.display = 'none'
        document.getElementById("single-order-container").style.display = 'none'
        document.getElementById("order-container").style.display = 'none'
        document.getElementById('single-order-container').style.display ='none';
        document.getElementById('checkout-container').style.display = 'none'
        document.getElementById("js-display-items").innerHTML = html;
        document.getElementById("js-display-items").style.display = 'block';
        HideHomePage()
        return
      }
      if (price >= 500) {
        html += `  <div class="row my-4">
        <div class="col-sm-6">
          <a onclick="HomePage()" class="btn btn-link text-muted">
            <i class="mdi mdi-arrow-left me-1"></i> Continue Shopping </a>
        </div>
        <div class="col-sm-6">
          <div class="text-sm-end mt-2 mt-sm-0">
            <a onclick="Checkout()" class="btn btn-success">
              <i class="mdi mdi-cart-outline me-1"></i> Checkout </a>
          </div>
        </div>
      </div>
    </div>
  
    <div class="col-xl-4">
      <div class="mt-5 mt-lg-0">
        <div class="card border shadow-none">
          <div class="card-header bg-transparent border-bottom py-3 px-4">
            <h5 class="font-size-16 mb-0">Order Summary <span class="float-end"></span></h5>
          </div>
          <div class="card-body p-4 pt-2">
            <div class="table-responsive">
              <table class="table mb-0">
                <tbody>
                  <tr>
                    <td>Sub Total :</td>
                    <td class="text-end">₹${price}</td>
                  </tr>
                  <tr>
                    <td>Shipping Charge :</td>
                    <td class="text-end">₹ FREE</td>
                  </tr>
                  <tr class="bg-light">
                    <th>Total :</th>
                    <td class="text-end">
                      <span class="fw-bold">
                      ₹${price}
                      </span>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
  
          </div>
        </div>
      </div>
    </div>
  </div>
  </div>
  </div>
  `
      } else {
        html += `  <div class="row my-4">
        <div class="col-sm-6">
          <a onclick="HomePage()" class="btn btn-link text-muted">
            <i class="mdi mdi-arrow-left me-1"></i> Continue Shopping </a>
        </div>
        <div class="col-sm-6">
          <div class="text-sm-end mt-2 mt-sm-0">
            <a onclick="Checkout()" class="btn btn-success" style="color:white;">
              <i class="mdi mdi-cart-outline me-1"></i> Checkout </a>
          </div>
        </div>
      </div>
    </div>
  
    <div class="col-xl-4">
      <div class="mt-5 mt-lg-0">
        <div class="card border shadow-none">
          <div class="card-header bg-transparent border-bottom py-3 px-4">
            <h5 class="font-size-16 mb-0">Order Summary <span class="float-end"></span></h5>
          </div>
          <div class="card-body p-4 pt-2">
            <div class="table-responsive">
              <table class="table mb-0">
                <tbody>
                  <tr>
                    <td>Sub Total :</td>
                    <td class="text-end">₹${price}</td>
                  </tr>
                  <tr>
                    <td>Shipping Charge :</td>
                    <td class="text-end">₹50</td>
                  </tr>
                  <tr class="bg-light">
                    <th>Total :</th>
                    <td class="text-end">
                      <span class="fw-bold">
                      ₹${price + 50}
                      </span>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
  
          </div>
        </div>
      </div>
    </div>
  </div>
  </div>
  </div>
  `

      }


      document.querySelector('.checkout-container').style.display = 'none'
      document.getElementById("single-order-container").style.display = 'none'
      document.getElementById("order-container").style.display ='none'
      document.getElementById('checkout-container').style.display = 'none'
      document.getElementById('single-order-container').style.display ='none';
      document.getElementById("js-display-items").innerHTML = html;
      document.getElementById("js-display-items").style.display = 'block';
      HideHomePage()

    })
    .catch(error => {
      showToast(error, "Error", 0);
    });
}

async function DeleteProduct(productName, productQuantity) {
  try {
    const storedData = localStorage.getItem("userdata");
    const retrievedUserData = JSON.parse(storedData);
    const data = {
      token: retrievedUserData.token,
      name: productName,
      quantity: Number(productQuantity) // Include quantity in the request data
    };


    const response = await fetch('http://localhost:8080/deleteproduct', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });

    const result = await response.json();
    if (result === true) {
      DisplayCart();
      showToast("Deleted Successfully", "Info", 3)
    } else {
      showToast("No Itmes found", "Info", 3)
    }
  } catch (error) {
    showToast('Error deleting product:' + error, "Danger", 0);
  }
}

async function IncrementButton(productName, productQuantity) {
  try {
    const storedData = localStorage.getItem("userdata");
    const retrievedUserData = JSON.parse(storedData);
    const data = {
      customerid: retrievedUserData.token,
      productname: productName,
      quantity: Number(productQuantity) + 1 // Include quantity in the request data
    };


    const response = await fetch('http://localhost:8080/updatecart', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });

    const result = await response.json();
    if (result.message) {
      DisplayCart();
      showToast(result.message, "Info", 3)
    } else if (result.error) {
      showToast(result.error, "Error", 0)
    }
  } catch (error) {
    showToast('Error deleting product:' + error, "Danger", 0);
  }
}


async function DecrementButton(productName, productQuantity) {
  try {
    const storedData = localStorage.getItem("userdata");
    const retrievedUserData = JSON.parse(storedData);

    const data = {
      customerid: retrievedUserData.token,
      productname: productName,
      quantity: Number(productQuantity) - 1 // Include quantity in the request data
    };


    const response = await fetch('http://localhost:8080/updatecart', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });

    const result = await response.json();
    if (result.message) {
      DisplayCart();
      showToast(result.message, "Info", 3)
    } else if (result.error) {
      showToast(result.error, "Error", 0)
    }
  } catch (error) {
    showToast('Error deleting product:' + error, "Danger", 0);
  }
}


async function Checkout() {
  try {
    GetUserAddress()

    const storedData = localStorage.getItem("userdata");
    const retrievedUserData = JSON.parse(storedData);

    const data = {
      token: retrievedUserData.token,
    };


    const response = await fetch('http://localhost:8080/products', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });
    const result = await response.json();
    console.log(result)
    let html = ""
    let price = 0
    if (result.message) {
      document.querySelector('.checkout-container').style.display = 'block';
      result.message.forEach((item) => {
        html += `
        <div class="d-flex align-items-center mb-4">
        <div class="me-3 position-relative">
          <span class="position-absolute top-0 start-100 translate-middle badge rounded-pill badge-secondary">
            ${item.quantity}
          </span>
          <img src="data:image/jpeg;base64,${item.image}"
            style="height: 96px; width: 96x;" class="img-sm rounded border" />
        </div>
        <div class="">
          <a href="#" class="nav-link">
            ${item.productname} <br />
            ${item.itemcategory}
          </a>
          <div class="price text-muted">Total: ₹${item.totalprice}</div>
        </div>
        </div>
      `;
        price += Number(item.totalprice)
      })
      if (price >= 500) {
        document.getElementById('itemscost').innerHTML = '₹' + price
        document.getElementById('shipingcost').innerHTML = '₹' + 0
        document.getElementById('totalcost').innerHTML = '₹' + price
      } else {
        document.getElementById('itemscost').innerHTML = '₹' + price
        document.getElementById('shipingcost').innerHTML = '₹' + 50
        document.getElementById('totalcost').innerHTML = '₹' + (price + 50)
      }
    } else if (result.error) {
      showToast(result.error, "Danger", 0);
      return
    }
    document.getElementById("single-order-container").style.display = 'none'
    document.getElementById("order-container").style.display ='none'
    document.getElementById('single-order-container').style.display ='none';
    document.getElementById('js-display-items').style.display = 'none';
    document.getElementById('checkout-product-container').innerHTML = html

  } catch (error) {
    console.log(error)
  }
}

// Get and Display User Address
async function GetUserAddress() {
  try {
    const storedData = localStorage.getItem("userdata");
    const retrievedUserData = JSON.parse(storedData);
    const data = {
      token: retrievedUserData.token,
    };
    const output = await fetch('http://localhost:8080/getuseraddress', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });
    const address = await output.json();
    if (address.message) {
      console.log(address.message)
      document.getElementById('address-firstname').value = address.message.firstname
      document.getElementById('address-lastname').value = address.message.lastname
      document.getElementById('address-email').value = address.message.deliveryemail
      document.getElementById('addess-streetname').value = address.message.streetname
      document.getElementById('address-pincode').value = address.message.pincode
      document.getElementById('address-city').value = address.message.city
      document.getElementById('address-phone').value = address.message.deliveryphoneno
    } else if (address.error) {
      showToast(address.error, "Danger", 0);
    }
  } catch (error) {
    showToast(error, "Error", 0)
  }
}


// Save User Address
async function SaveAddress() {
  try {
    const storedData = localStorage.getItem("userdata");
    const retrievedUserData = JSON.parse(storedData);
    const data = {
      token: retrievedUserData.token,
      firstname: document.getElementById('address-firstname').value,
      lastname: document.getElementById('address-lastname').value,
      deliveryemail: document.getElementById('address-email').value,
      deliveryphoneno: Number(document.getElementById('address-phone').value),
      streetname: document.getElementById('addess-streetname').value,
      city: document.getElementById('address-city').value,
      pincode: Number(document.getElementById('address-pincode').value),
    }
    console.log(data)
    const output = await fetch('http://localhost:8080/adddeliveryaddress', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });
    const address = await output.json();
    if (address.message) {
      showToast(address.message, "Info", 3)
    } else if (address.error) {
      showToast(address.error, "Error", 0)
    }

  } catch (error) {
    showToast(error, "Error", 0)
  }
}

async function PayNow() {
  try {
    SaveAddress()
    const storedData = localStorage.getItem("userdata");
    const retrievedUserData = JSON.parse(storedData);
    const data = {
      token: retrievedUserData.token,
    }
    const output = await fetch('http://localhost:8080/buynow', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });
    const value = await output.json();
    if (value.message) {
      showToast("Your order has been Placed", "Success", 3)
      DisplayPlacedImage()
    } else if (value.error) {
      showToast(value.error, "Error", 0)
    }
  } catch {
    console.log(error)
  }
}

async function DisplayPlacedImage() {
  try {
    document.querySelector('.checkout-container').style.display = 'none'
    document.getElementById("single-order-container").style.display = 'none'
    document.getElementById('single-order-container').style.display ='none';
    document.getElementById("order-container").style.display ='none'
    HideHomePage()
    document.querySelector('.checkout-container').style.display = 'none';
    document.getElementById('js-display-items').innerHTML =
      `
            <div class="col-sm-6">
            <a onclick="HomePage()" class="btn btn-link text-muted">
              <i class="mdi mdi-arrow-left me-1"></i> Continue Shopping </a>
            </div>   
            <div class="order-conformation=image" style="margin-right:30px">
              <center>
                <img class="conformation-image" src="./images/output-onlinegiftools.gif">
              </center>
            </div>
           `
    document.getElementById('js-display-items').style.display = 'block';
  } catch {
    console.log(error)
  }
}


async function DisplayOrders() {
  try {
    const storedData = localStorage.getItem("userdata");
    const retrievedUserData = JSON.parse(storedData);
    const data = {
      token: retrievedUserData.token,
    }
    console.log(data)
    const output = await fetch('http://localhost:8080/customerorders', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });
    const value = await output.json();
    if (value.message) {

      html = ""
      value.message.forEach((item) => {
        console.log(item)
        html += `           
        <div class="container-fluid my-5 justify-content-center"
        style="margin-top: 20px !important; margin-bottom: 50px !important; display: flex; justify-content: center; align-items: center; flex-direction: column;">
        <div class="card-container" style="width: 1000px; margin-bottom: 5px;">
          <div class="card px-2"
            style="background-color: #fff; box-shadow: 1px 2px 5px 0px rgb(0, 0, 0); z-index: 0; padding: 5px; margin-bottom: -30px;">
            <div class="card-header bg-white">
              <div class="row justify-content-between">
                <div class="col">
                  <p class="text-muted"> Order ID <span class="font-weight-bold text-dark">${item.orderid}</span></p>
                  <p class="text-muted"> Ordred On <span class="font-weight-bold text-dark">${item.orderdate}</span>
                  </p>
                </div>
              </div>
            </div>
            <div class="card-body" style="padding: 10px;">
              <div class="media flex-column flex-sm-row">
                <div class="media-body">
                  <h5 class="bold">${item.itemstobuy.productname}</h5>
                  <p style="margin-bottom: 5px;"> Qt: ${item.noofitems}</p> <!-- Reduced margin -->
                  <h4 style="margin-top: -30px; margin-bottom: 10px;"> <span class="mt-5">&#x20B9;</span> ${item.itemstobuy.totalprice}<span
                      style="font-size: medium;">  &nbsp;via (COD) </span></h4>
                  <p class="text-muted">Estimated Delivery Date: <span class="Today">${item.deliverydate}</span></p>
                </div>
                <img class="align-self-center img-fluid" src="data:image/jpeg;base64,${item.itemstobuy.image}" width="150" height="150">
              </div>
            </div>
  
            <div class="card-footer bg-white px-sm-3 pt-sm-4 px-0" style="padding: 5px;">
              <div class="row text-center">
                <div class="col my-auto border-line" style=" border-right: 1px solid rgb(226, 206, 226);cursor:pointer" onclick="GetOrder('${item.orderid}')">
                  <h5 >View Details</h5>
                </div>
                <div class="col my-auto border-line" onclick="DeleteOrder('${item.orderid}')" style="cursor:pointer;">
                  <h5>Cancel Order</h5>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
  </div>`
      })
      if (html == "") {
        html = `<img src="./images/emptyorder.gif">`
      }
      HideHomePage()
      document.querySelector('.checkout-container').style.display = 'none'
      document.getElementById("single-order-container").style.display = 'none'
      document.getElementById("order-container").style.display ='block'
      document.getElementById('checkout-container').style.display = 'none'
      document.getElementById('single-order-container').style.display ='none';
      document.getElementById("js-display-items").style.display = 'none';
      document.getElementById('display-order-conatiner').innerHTML = html;


    } else if (value.error) {
      showToast(error, "Error", 0)
    }
  } catch (error) {
    showToast(error, "Error", 0)
  }
}

async function GetOrder(id) {
  try {
    console.log("On GetOrder")
    const storedData = localStorage.getItem("userdata");
    const retrievedUserData = JSON.parse(storedData);
    const data = {
      token: retrievedUserData.token,
      orderid: id,
    }
    console.log(data)
    const output = await fetch('http://localhost:8080/getcustomerorder', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });
    const value = await output.json();
    if (value.message) {

      let html = ""
      let orderstatus = ""
      const item = value.message
      const entries = Object.entries(item.status);

      entries.forEach(([key, val]) => {
        if (val == "completed") {
          orderstatus = key
          return
        }
      });
      html = ` <div class="container">

        <div class="d-flex justify-content-between align-items-center py-3">
        </div>
  
        <div class="row">
          <div class="col-lg-8">
  
            <div class="card mb-4">
              <div class="card-body">
                <div class="mb-3 d-flex justify-content-between">
                  <div>
                    <span class="me-3">${item.orderdate}</span>
                    <span class="me-3">${item.orderid}</span>
                    <!-- <span class="me-3">Visa -1234</span> -->
                    <span class="badge rounded-pill bg-info">${orderstatus.toUpperCase()}</span>
                  </div>
                </div>
                <table class="table table-borderless">
                  <tbody>
                    <tr>
                      <td>
                        <div class="d-flex mb-2">
                          <div class="flex-shrink-0">
                            <img src="data:image/jpeg;base64,${item.itemstobuy.image}" alt width="35"
                              class="img-fluid">
                          </div>
                          <div class="flex-lg-grow-1 ms-3">
                            <h6 class="small mb-0"><a href="#" class="text-reset">${item.itemstobuy.productname}</a></h6>
                            <span class="small">Category: ${item.itemstobuy.itemcategory}</span>
                          </div>
                        </div>
                      </td>
                      <td>${item.noofitems}</td>
                      <td class="text-end">₹${item.itemstobuy.price}</td>
                    </tr>
  
                  </tbody>
                  <tfoot>
                    <tr class="fw-bold">
                      <td colspan="2">TOTAL</td>
                      <td class="text-end">₹${item.totalamount}</td>
                    </tr>
                  </tfoot>
                </table>
              </div>
            </div>
  
  
            <div class="card-body">
              <div class="row">
  
                <div class="card mb-3">
                  <div class="d-flex flex-wrap flex-sm-nowrap justify-content-between py-3 px-2 bg-secondary">
                    <div class="w-100 text-center py-1 px-2"><span class="text-medium">Shipped By:</span> ${item.itemstobuy.sellerid}
                    </div>
                    <div class="w-100 text-center py-1 px-2"><span class="text-medium">Status:</span> ${orderstatus.toUpperCase()}
                    </div>
                    <div class="w-100 text-center py-1 px-2"><span class="text-medium">Expected Date:</span>${item.deliverydate}
                    </div>
                  </div>
                  <div class="card-body">
                    <div
                      class="steps d-flex flex-wrap flex-sm-nowrap justify-content-between padding-top-2x padding-bottom-1x">
                      <div class="step ${item.status.confirmed}">
                        <div class="step-icon-wrap">
                          <div class="step-icon"><i class="pe-7s-cart"></i></div>
                        </div>
                        <h4 class="step-title">Confirmed Order</h4>
                      </div>
                      <div class="step ${item.status.processing}">
                        <div class="step-icon-wrap">
                          <div class="step-icon"><i class="pe-7s-config"></i></div>
                        </div>
                        <h4 class="step-title">Processing Order</h4>
                      </div>
                      <div class="step ${item.status.quality}">
                        <div class="step-icon-wrap">
                          <div class="step-icon"><i class="pe-7s-medal"></i></div>
                        </div>
                        <h4 class="step-title">Quality Check</h4>
                      </div>
                      <div class="step ${item.status.dispatched}">
                        <div class="step-icon-wrap">
                          <div class="step-icon"><i class="pe-7s-car"></i></div>
                        </div>
                        <h4 class="step-title">Product Dispatched</h4>
                      </div>
                      <div class="step ${item.status.delivered}">
                        <div class="step-icon-wrap">
                          <div class="step-icon"><i class="pe-7s-home"></i></div>
                        </div>
                        <h4 class="step-title">Product Delivered</h4>
                      </div>
                    </div>
                  </div>
                </div>
                <div
                  class="d-flex flex-wrap flex-md-nowrap justify-content-center justify-content-sm-between align-items-center">
                  <div class="custom-control custom-checkbox mr-3">
                  </div>
                </div>
  
  
  
              </div>
            </div>
          </div>
          <div class="col-lg-4">
  
            <div class="card mb-4">
              <div class="card-body">
                <h3 class="h6">Customer Notes</h3>
                <p>The package will be delivered to this below mentioned address only.</p><p>If there is any problem feel free to contact us</p>
              </div>
            </div>
            <div class="card mb-4">
  
              <div class="card-body">
                <h3 class="h6">Shipping Information</h3>
                <strong>Phone:</strong>
                <span><a href="#" class="text-decoration-underline" target="_blank">${item.address.deliveryphoneno}</a> <i
                    class="bi bi-box-arrow-up-right"></i> </span>
                <hr>
                <h3 class="h6">Address</h3>
                <address>
                  <strong>${item.address.firstname} ${item.address.lastname}</strong><br>
                    ${item.address.streetname}<br>
                    ${item.address.city} - ${item.address.pincode}
                  <br>
                  ${item.address.deliveryemail}
                </address>
              </div>
            </div>
          </div>
        </div>
      </div>`
      document.querySelector('.checkout-container').style.display = 'none'
      document.getElementById("order-container").style.display ='none'
      document.getElementById('checkout-container').style.display = 'none'
      document.getElementById("js-display-items").style.display = 'none';
      HideHomePage()
      document.getElementById('single-order-container').innerHTML = html;
      document.getElementById('single-order-container').style.display ='block';

    } else if (value.error) {
      showToast(value.error, "Error", 0)
    }

  } catch (error) {
    console.log(error)
  }
}

async function DeleteOrder(orderid){
  try{
    const storedData = localStorage.getItem("userdata");
    const retrievedUserData = JSON.parse(storedData);
    const data = {
      token: retrievedUserData.token,
      orderid: orderid,
    }
    console.log(data)
    const output = await fetch('http://localhost:8080/deleteorder', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    });
    const value = await output.json();
    if(value.message){
      showToast(value.message,"Success",3)
      DisplayOrders()
      return
    }else if(value.error){
      showToast(value.message,"Error",3)
    }

  }catch(error){
     console.log(error)
  }

}





// document.querySelector('#cart-form').addEventListener('submit', function (event) {
//     event.preventDefault(); // Prevent the default form submission behavior

//     const titleElement = document.querySelector('.showcase-title1');
//     const priceElement = document.querySelector('.pricesoap'); // Note the class name change
//     const itemTitle = titleElement.textContent;
//     const priceText = priceElement.textContent;
//     const price = parseFloat(priceText.replace('$', ''));
//     const dataToSend = {
//         name: itemTitle,
//         price: price,
//     };

//     fetch('http://localhost:8080/addtocart', {
//         method: 'POST',
//         headers: {
//             'Content-Type': 'application/json',
//         },
//         body: JSON.stringify(dataToSend),
//     })
//         .then((response) => response.json())
//         .then((data) => {
//             console.log('Item Title stored in the database:', data);
//         })
//         .catch((error) => {
//             console.error('Error storing item title:', error);
//         });
// });

// const itemLinks = document.querySelectorAll(".item-link");
// const baseURL = "/items"; // The base URL for your items route

// itemLinks.forEach(link => {
//   link.addEventListener("click", function (event) {
//     event.preventDefault();
//     const storedData = localStorage.getItem("userdata");
//     const retrievedUserData = JSON.parse(storedData);
//     const itemName = link.getAttribute("data-item"); // Get the item name
//     const token = urlParams.get('token') || retrievedUserData.token; // Replace with your dynamic token logic

//     const url = `${baseURL}?item=${itemName}&token=${token}`;
//     window.location.href = url; // Redirect to the dynamic URL
//   });
// });

// document.addEventListener("DOMContentLoaded", function () {
//     const token1 = token; // Replace with your actual token value
//     const searchBtn = document.getElementById("personoutline");
//     searchBtn.addEventListener("click", function () {
//       const url = `/ordereditems?token=${token1}`;
//       window.location.href = url;
//     });
//   });

//   document.addEventListener("DOMContentLoaded", function () {
//     const token1 = token; // Replace with your actual token value
//     const searchBtn = document.getElementById("searchBtn");
//     searchBtn.addEventListener("click", function () {
//       const url = `/inventory/?token=${token1}`;
//       window.location.href = url;
//     });
//   });

//   var urlParams = new URLSearchParams(window.location.search);

//   var token = urlParams.get('token');

//   document.addEventListener("DOMContentLoaded", function () {
//     var cartButton = document.querySelector(".action-btn a");
//     if (cartButton) {
//       cartButton.href = "/cart/?token=" + token;
//     } else {
//       console.error("Button not found in the DOM.");
//     }
//   })
