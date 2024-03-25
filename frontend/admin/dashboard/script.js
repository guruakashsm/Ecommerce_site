var adminData = localStorage.getItem('admindata');
if (adminData) {
    // Parse the JSON string to convert it into a JavaScript object
    var adminObject = JSON.parse(adminData);
    document.querySelector('.admin-name').innerHTML = ExtractUserName(adminObject.username)
    document.querySelector('.admin-target').innerHTML = `Congratulations ${ExtractUserName(adminObject.username)}`
} else {
    showToast('Please Login as admin', "Danger", 0);
}
const formData = {
    "token": adminData.token
}



function DisplayData() {
    fetch("http://localhost:8080/adminpage", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
    })
        .then(response => response.json())
        .then(data => {
            if (data.result) {
                document.querySelector('.customer-count').innerHTML = `${data.result.usercount}k`
                document.querySelector('.seller-count').innerHTML = `${data.result.sellercount}k`
                document.querySelector('.product-count').innerHTML = `${data.result.productcount}k`
                document.querySelector('.sales-count').innerHTML = `${data.result.salescount}k`


            }

        })
        .catch(error => {
            showToast(error.message, "Danger", 0);
        });
}
DisplayData()



function ExtractUserName(name) {
    var adminname = "";
    for (let i = 0; i < name.length; i++) {
        if (name[i] == "@") {
            return adminname.toUpperCase();
        } else {
            adminname += name[i];
        }
    }
    return adminname.toUpperCase();
}


function LogOut() {
    localStorage.removeItem("admindata")
    localStorage.removeItem("adminsignindata")
    window.location.href = "/ecom/adminlogin"
}

function DisplaySellerData() {
    let count = 0
    fetch("http://localhost:8080/getallsellerdata", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
    })

        .then(response => response.json())
        .then(data => {
            let html = ""
            if (data.seller) {

                data.seller.forEach((element, index) => {
                    count++
                    if (count == 5) {
                        document.querySelector('.js-seller-data').innerHTML = html
                        return
                    }
                    html += `<li class="d-flex mb-4 pb-md-2">
                    <div class="d-flex w-100 flex-wrap align-items-center justify-content-between gap-2">
                      <div class="me-2">
                        <h6 class="mb-0">${element.sellername.toUpperCase()}</h6>
                        <small>${element.selleremail}</small>
                      </div>
                      <div>
                        <h6 class="mb-2">${element.phoneno}</h6>
                    </div>
                  </li>`
                })
                document.querySelector('.js-seller-data').innerHTML = html
                return
            }

        })
        .catch(error => {
            showToast(error.message, "Danger", 0);
        });
}
DisplaySellerData();

function DisplayWorkers() {
    fetch("http://localhost:8080/getworkers", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
    })
        .then(response => response.json())
        .then(data => {
            let html = ""
            if (data.result) {
                data.result.forEach((element, index) => {
                    html += `<tr>
                    <td>
                      <div class="d-flex align-items-center">
                        <div class="avatar avatar-sm me-3">
                          <img src="data:image/jpeg;base64,${element.image}" alt="Avatar" class="rounded-circle" />
                        </div>
                        <div>
                          <h6 class="mb-0 text-truncate">${element.username.toUpperCase()}</h6>
                        </div>
                      </div>
                    </td>
                    <td class="text-truncate">${element.email}</td>
                    <td class="text-truncate">
                      <i class="mdi mdi-laptop mdi-24px text-danger me-1"></i>${element.role}
                    </td>
                    <td class="text-truncate">${element.no}</td>
                    <td class="text-truncate">${element.salary}</td>
                    <td><span class="badge bg-label-warning rounded-pill">${element.status}</span></td>
                  </tr>`
                })
                document.querySelector('.js-workers').innerHTML = html

            }

        })
        .catch(error => {
            showToast(error.message, "Danger", 0);
        });
}
DisplayWorkers()


function TotlalSales() {
    fetch("http://localhost:8080/adminpage", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
    })
        .then(response => response.json())
        .then(data => {
            if (data.result) {
                console.log(data.result)
                document.querySelector('.customer-count').innerHTML = `${data.result.usercount}k`
                document.querySelector('.seller-count').innerHTML = `${data.result.sellercount}k`
                document.querySelector('.product-count').innerHTML = `${data.result.productcount}k`
                document.querySelector('.sales-count').innerHTML = `${data.result.salescount}k`
                document.querySelector('.total-growth').innerHTML = `Total Revenue Gained : &#8377; ${data.result.totalsalesamount}`
                document.querySelector('.total-sales-profit').innerHTML = ` &#8377; ${data.result.totalsalesamount}`
                document.querySelector('.target-persentage').innerHTML = `${calculatePercentage(10000, data.result.totalsalesamount)}% of target 🚀`
                document.querySelector('.profit').innerHTML = `&#8377;${data.result.totalsalesamount}`
                document.querySelector('.profit-persent').innerHTML = `+${calculatePercentage(1000, data.result.totalsalesamount)}%`




            }

        })
        .catch(error => {
            showToast(error.message, "Danger", 0);
        });
}

TotlalSales();

function calculatePercentage(totalAmount, receivedAmount) {
    if (totalAmount <= 0) {
        showToast("Total amount should be greater than zero.", "Info", 1);
        return null;
    }

    const percentage = (receivedAmount / totalAmount) * 100;
    return percentage.toFixed(2); // Round to two decimal places
}
function DisplayFeedBack() {
    let count = 0;
    fetch("http://localhost:8080/getfeedback", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
    })
        .then(response => response.json())
        .then(data => {
            let html = ""
            if (data.result) {

                data.result.forEach((element, index) => {
                    count++
                    if (count == 6) {
                        document.querySelector('.js-seller-data').innerHTML = html
                        return
                    }
                    element.role = element.role.toUpperCase()
                    html += `<li class="d-flex mb-4 pb-md-2">
                    <div class="d-flex w-100 flex-wrap align-items-center justify-content-between gap-2">
                      <div class="me-2">
                        <h6 class="mb-0">${element.email}</h6>
                        <small>${element.feedback}</small>
                      </div>
                      <div class="d-flex align-items-center">
                        <h9 class="mb-2">${element.role}</h9>
                        
                      </div>
                      <img src="./images/success.png" height="17px" style="cursor:pointer" alt="Delete" class="delete-icon" onclick="deleteFeedback('${element.email}','${element.feedback}')">
                    </div>
                  </li>`
                })
                document.querySelector('.js-feedback-data').innerHTML = html

            }

        })
        .catch(error => {
            showToast(error.message, "Danger", 0);
        });
}
DisplayFeedBack()

function deleteFeedback(email, feedback) {
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
                showToast("FeedBack Deleted Successfully", "Success", 3);
                DisplayFeedBack()
            } else {
                showToast("Error deleting feedback", "Danger", 0);
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


function PrintContent() {

    var printableElement = document.getElementById('printableContent');

    // Open a new window for printing
    var printWindow = window.open('', '_blank');

    // Write the printable content to the new window
    printWindow.document.write('<html><head><title>Print</title></head><body>');
    printWindow.document.write(printableElement.innerHTML);
    printWindow.document.write('</body></html>');

    // Close the document stream and trigger the print dialog
    printWindow.document.close();
    printWindow.print();
}


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
                    showToast("FeedBack Deleted Successfully", "Success", 3);
                    feedbackBox.remove();
                } else {
                    showToast("Error deleting feedback", "Danger", 0);
                }
            })
            .catch(error => {
                showToast(error, "Error", 0);
            });
    });
    feedbackBox.appendChild(deleteIcon);
}

function DisplayListUsers() {
    console.log("Displaylist")
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.querySelector('.outer-container').style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
    document.getElementById('Inventorysnip').style.display = 'none';
    document.querySelector('.wrapper').style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    document.getElementById("notapprovedseller").style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById('update-form-admin-container').style.display = 'none';
    document.getElementById('snippetContent').style.display = 'block';
    fetch('http://localhost:8080/getallcustomerdata')
        .then(response => response.json())
        .then(data => {
            let html = ""

            data.forEach(customer => {

                html += `
            <tr class="candidates-list customer-list">
            <td class="title">
              <div class="thumb"> <img class="img-fluid"
                  src="http://previews.123rf.com/images/jenjawin/jenjawin1904/jenjawin190400251/120265520-account-icon-outline-vector-eps10-user-profile-sign-web-icon-with-check-mark-glyph-user-authorized.jpg" alt="">
              </div>
              <div class="candidate-list-details">
                <div class="candidate-list-info">
                  <div class="candidate-list-title customer">
                    <h5 class="mb-0"><a href="#">${customer.name.toUpperCase()}</a></h5>
                  </div>
                  <div class="candidate-list-option">
                    <ul class="list-unstyled">
                      <li><i class="fas fa-filter pr-1"></i>${customer.email}
                      </li>
                      <li><i class="fas fa-map-marker-alt pr-1"></i>${customer.address}</li>
                    </ul>
                  </div>
                </div>
              </div>
            </td>
            <td class="candidate-list-favourite-time text-center"> <a
                class="candidate-list-favourite order-2 text-danger" href="#"></a>
              <span class="candidate-list-time order-1">${customer.phonenumber}</span></td>
            <td>
              <ul class="list-unstyled mb-0 d-flex justify-content-end">
              <li onclick="ViewData('${customer.email}','customer');recentPage = 'customer';"><a  class="text-info" data-toggle="tooltip" title="" data-original-title="Edit"><i
              class="fas fa-eye"></i></a>
              </li>
              <li onclick="EditData('${customer.email}','customer');recentPage = 'customer';"><a class="text-info" data-toggle="tooltip" title="" data-original-title="Edit"><i
              class="fas fa-pencil-alt"></i></a>
              </li>
                <li  onclick="DeleteData('${customer.email}','cus');recentPage = 'customer';"><a class="text-danger" data-toggle="tooltip" title=""
                    data-original-title="Delete"><i class="far fa-trash-alt"></i></a></li>
                <li  onclick="BlockUser('${customer.email}','customer');DisplayListUsers();recentPage = 'customer';"><a class="text-danger" data-toggle="tooltip" title=""
                    data-original-title="Delete"><i class="fas fa-ban"></i></a></li>
              </ul>
            </td>
          </tr>`;

            });
            document.querySelector('.user-list-body').innerHTML = html;
        })
        .catch(error => {
            showToast(error, "Error", 0);
        });
}
function DeleteData(email, coll) {
    const requestData = {
        collection: "",
        idValue: email
    };
    if (coll == 'cus') {
        const collection = "customer";
        requestData.collection = collection
    } else if (coll == "sel" || coll == "seller") {
        const collection = "seller";
        requestData.collection = collection
    } else if (coll == "inven") {
        const collection = "inventory";
        requestData.collection = collection
    }



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
            ;
            if (data === true) {
                showToast("Deleted Sucessfull", "Success", 3)
                if (coll == 'cus') {
                    DisplayListUsers()
                } else if (coll == "sel") {
                    DisplayListSeller()
                } else if (coll == "inven") {
                    DisplayListInventory()
                } else if (coll == "seller") {
                    Displayallnotapprovedseller()
                }

            } else {
                showToast("Error in Deleting", "Danger", 0)
            }
        })
        .catch(error => {
            showToast(error.message, "Error", 0);
        });


}


function DisplayListSeller() {
    console.log("Displaylist")
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.getElementById('snippetContent').style.display = 'none';
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.querySelector('.outer-container').style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    document.querySelector('.wrapper').style.display = 'none';
    document.getElementById('Inventorysnip').style.display = 'none';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'none';
    document.getElementById("notapprovedseller").style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.getElementById('update-form-admin-container').style.display = 'none';
    document.getElementById('sellersnip').style.display = 'block';
    fetch("http://localhost:8080/getallsellerdata", {
        method: "POST", // Use DELETE method to delete data
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify()
    })

        .then(response => response.json())
        .then(data => {
            let html = ""

            data.seller.forEach(seller => {

                html += `
            <tr class="candidates-list seller-list">
            <td class="title">
              <div class="thumb"> <img class="img-fluid"
                  src="data:image/jpeg;base64,${seller.image}" alt="">
              </div>
              <div class="candidate-list-details">
                <div class="candidate-list-info">
                  <div class="candidate-list-title seller">
                    <h5 class="mb-0"><a href="#">${seller.sellername.toUpperCase()}</a></h5>
                  </div>
                  <div class="candidate-list-option">
                    <ul class="list-unstyled">
                      <li><i class="fas fa-filter pr-1"></i>${seller.selleremail}
                      </li>
                      <li><i class="fas fa-map-marker-alt pr-1"></i>${seller.address}</li>
                    </ul>
                  </div>
                </div>
              </div>
            </td>
            <td class="candidate-list-favourite-time text-center"> <a
                class="candidate-list-favourite order-2 text-danger" href="#"></a>
              <span class="candidate-list-time order-1">${seller.phoneno}</span></td>
            <td>
              <ul class="list-unstyled mb-0 d-flex justify-content-end">
              <li onclick="ViewData('${seller.selleremail}','seller');recentPage = 'inventory';"><a  class="text-info" data-toggle="tooltip" title="" data-original-title="Edit"><i
              class="fas fa-eye"></i></a>
              </li>
              <li onclick="EditData('${seller.selleremail}','seller');recentPage = 'inventory';"><a class="text-info" data-toggle="tooltip" title="" data-original-title="Edit"><i
              class="fas fa-pencil-alt"></i></a>
              </li>
                <li  onclick="DeleteData('${seller.selleremail}','sel');recentPage = 'inventory';"><a class="text-danger" data-toggle="tooltip" title=""
                    data-original-title="Delete"><i class="far fa-trash-alt"></i></a></li>
                    <li  onclick="BlockUser('${seller.selleremail}','seller');DisplayListSeller();recentPage = 'seller';"><a class="text-danger" data-toggle="tooltip" title=""
                data-original-title="Delete"><i class="fas fa-ban"></i></a></li>
              </ul>
            </td>
          </tr>`;

            });
            document.querySelector('.seller-list-body').innerHTML = html;
        })
        .catch(error => {
            showToast(error, "Error", 0);
        });
}

function DisplayListInventory() {
    console.log("Displaylist")
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.getElementById("notapprovedseller").style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.querySelector('.outer-container').style.display = 'none';
    document.querySelector('.wrapper').style.display = 'none';
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.getElementById('Inventorysnip').style.display = 'block';
    document.getElementById('update-form-admin-container').style.display = 'none';
    fetch("http://localhost:8080/getallinventorydata", {
        method: "GET", // Use DELETE method to delete data
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify()
    })

        .then(response => response.json())
        .then(data => {
            let html = ""

            data.Inventory.forEach(customer => {

                html += `
            <tr class="candidates-list inventory-list">
            <td class="title">
              <div class="thumb"> <img class="img-fluid"
              src="data:image/jpeg;base64,${customer.image}" alt="">
              </div>
              <div class="candidate-list-details">
                <div class="candidate-list-info">
                  <div class="candidate-list-title inventory">
                    <h5 class="mb-0"><a href="#">${customer.itemname.toUpperCase()}</a></h5>
                  </div>
                  <div class="candidate-list-option">
                    <ul class="list-unstyled">
                      <li><i class="fas fa-filter pr-1"></i>${customer.itemcategory.toUpperCase()}
                      </li>
                      <li><i class="fas fa-map-marker-alt pr-1"></i>${customer.quantity}</li>
                    </ul>
                  </div>
                </div>
              </div>
            </td>
            <td class="candidate-list-favourite-time text-center"> <a
                class="candidate-list-favourite order-2 text-danger" href="#"></a>
              <span class="candidate-list-time order-1">${customer.price}</span></td>
            <td>
              <ul class="list-unstyled mb-0 d-flex justify-content-end">
              <li onclick="ViewData('${customer.itemname}','inventory');recentPage = 'inventory';"><a  class="text-info" data-toggle="tooltip" title="" data-original-title="Edit"><i
              class="fas fa-eye"></i></a>
              </li>
              <li onclick="EditData('${customer.itemname}','inventory');recentPage = 'inventory';"><a  class="text-info" data-toggle="tooltip" title="" data-original-title="Edit"><i
              class="fas fa-pencil-alt"></i></a>
              </li>
                <li  onclick="DeleteData('${customer.itemname}','inven');recentPage = 'inventory';"><a class="text-danger" data-toggle="tooltip" title=""
                    data-original-title="Delete"><i class="far fa-trash-alt"></i></a></li>
              </ul>
            </td>
          </tr>`;

            });
            document.querySelector('.inventory-list-body').innerHTML = html;
        })
        .catch(error => {
            showToast(error, "Error", 0);
        });
}

function EditData(id, coll) {
    document.getElementById("updatecollection").value = coll;
    document.getElementById("idname").value = id;
    populateFieldOptions();
    DisplayEdit();
}



function Deletedata() {
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.querySelector('.outer-container').style.display = 'block';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.getElementById("notapprovedseller").style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.getElementById('Inventorysnip').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.querySelector('.wrapper').style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.getElementById('update-form-admin-container').style.display = 'none';
}

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
                showToast("Deleted Sucessfull", "Success", 3)
                document.getElementById("id").value = "";
            } else {
                showToast("Error in Deleting", "Danger", 0)
            }
        })
        .catch(error => {
            const resultDiv = document.getElementById("result-container");
            resultDiv.innerHTML = `<p>Error: ${error.message}</p>`;
        });
});

function DisplayEdit() {
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById("notapprovedseller").style.display = 'none';
    document.querySelector('.outer-container').style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.getElementById('Inventorysnip').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.querySelector('.wrapper').style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.getElementById('update-form-admin-container').style.display = 'block';
}

const updateFormElement = document.getElementById("update-form-admin");
const collectionSelectElement = document.getElementById("updatecollection");
const fieldSelectElement = document.getElementById("field");

const collectionselectOptions = {
    customer: ["name", "email", "phonenumber", "age", "password", "firstname", "lastname", "houseno", "streetname", "city", "pincode"],
    inventory: ["itemcategory", "itemname", "price", "quantity"],
    seller: ["sellername", "selleremail", "password", "phoneno", "address"],
};
function capitalizeFirstLetter(str) {
    return str.charAt(0).toUpperCase() + str.slice(1);
}

function populateFieldOptions() {
    const selectedCollection = collectionSelectElement.value || customer;
    const options = collectionselectOptions[selectedCollection] || [];

    // Clear existing options
    fieldSelectElement.innerHTML = "";

    // Add new options
    options.forEach(option => {
        const optionElement = document.createElement("option");
        optionElement.value = option;
        optionElement.textContent = capitalizeFirstLetter(option);
        fieldSelectElement.appendChild(optionElement);
    });
}
populateFieldOptions()



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
    console.log(requestData)

    fetch("http://localhost:8080/update", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(requestData)
    })
        .then(response => response.json())
        .then(data => {

            if (data) {
                showToast("Updated Successfully", "Success", 3)
                document.getElementById("update-form").reset();
            } else {
                showToast("Update Failes", "Danger", 0)
                document.getElementById("update-form").reset();
            }
        })
        .catch(error => {
            const resultDiv = document.getElementById("result");
            resultDiv.innerHTML = `<p>Error: ${error.message}</p>`;
        });
});


function CreateSeller() {
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    document.querySelector('.wrapper').style.display = 'block';
    document.getElementById("notapprovedseller").style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
    document.querySelector('.outer-container').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById('Inventorysnip').style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.getElementById('update-form-admin-container').style.display = 'none';
}

const sellerForm = document.getElementById('seller-form');
sellerForm.addEventListener('submit', (e) => {
    e.preventDefault();
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
        fetch('http://localhost:8080/createseller', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(sellerData)
        })
            .then(response => response.json())
            .then(data => {
                showToast('Seller created successfully.', 'Success', 3);
                sellerForm.reset();
            })
            .catch(error => {
                showToast(error, 'Error', 0);
            });
    };

    // Read the image file as a data URL
    reader.readAsArrayBuffer(imageFile);
});




function DisplayDrashBord() {
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.querySelector('.wrapper').style.display = 'none';
    document.querySelector('.outer-container').style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'block';
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
    document.getElementById("notapprovedseller").style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.getElementById('Inventorysnip').style.display = 'none';
    document.getElementById('update-form-admin-container').style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'none';
}

function CreateWorker() {
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.querySelector('.wrapper').style.display = 'none';
    document.querySelector('.outer-container').style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    document.getElementById("notapprovedseller").style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.getElementById('Inventorysnip').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.getElementById('feedbacksnip').style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
    document.getElementById('update-form-admin-container').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'block';
    document.getElementById('admin-wrapper').style.display = 'none';
}
document.getElementById("employee-wrapper").addEventListener("submit", function (event) {
    event.preventDefault();
    const imageFile = document.getElementById("worker-image").files[0];
    const reader = new FileReader();

    reader.onload = function () {
        const base64Image = btoa(new Uint8Array(reader.result).reduce((data, byte) => data + String.fromCharCode(byte), ''));
        const formData = {
            username: document.getElementById('worker-name').value,
            email: document.getElementById('worker-email').value,
            role: document.getElementById('worker-role').value,
            no: document.getElementById('worker-phone').value,
            salary: parseInt(document.getElementById('worker-salary').value),
            status: "Active",
            image: base64Image,
        }
        console.log(formData)

        // Send a POST request to your Go backend
        fetch("http://localhost:8080/createworker", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(formData),
        })
            .then(response => response.json())
            .then(data => {
                showToast(data.result, "info", 1)
            })
            .catch(error => {
                // Handle errors, e.g., display an error message
                showToast(error.message, "Error", 0);
            });
    };

    reader.readAsArrayBuffer(imageFile);
});

function DisplayCreateAdmin() {
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    document.querySelector('.wrapper').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.querySelector('.outer-container').style.display = 'none';
    document.getElementById("notapprovedseller").style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById('Inventorysnip').style.display = 'none';
    document.getElementById('update-form-admin-container').style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'block';
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.getElementById("qr-code").style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.getElementById("admin-input").style.display = 'block';
}

document.getElementById("admin-wrapper").addEventListener("submit", function (event) {
    event.preventDefault();

    const formData = {
        name: document.getElementById('admin-name').value,
        email: document.getElementById('admin-email').value,
        password: document.getElementById('admin-password').value,
        ip: document.getElementById('admin-ip').value,
        confirmpassword: document.getElementById('admin-confirmpassword').value
    }
    if (formData.email == "" || formData.password == "" || formData.ip == "") {
        showToast("Please Enter all Feilds", "Info", 1)
        return
    }
    if (formData.confirmpassword != formData.password) {
        showToast("Passoword Mismatch", "Danger", 0)
        return
    }
    console.log(formData)

    // Send a POST request to your Go backend
    fetch("http://localhost:8080/createadmin", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
    })
        .then(response => response.json())
        .then(data => {
            if (data.error != "") {
                showToast(data.error, "info", 1)
                return
            }
            showToast("Created Successfully", "info", 1)
            var totpKey = data.result;

            // Generate QR code
            var qrcode = new QRCode(document.getElementById("qrcode"), {
                text: totpKey,
                width: 128,
                height: 128,
            });
            document.getElementById("admin-form").reset()
            document.getElementById("qr-code").style.display = 'block';
            document.getElementById("admin-input").style.display = 'none';
            document.querySelector('.totp').innerHTML = `<br> Please Scan this QR to Get TOTP (or) <br> Use this Key : ${totpKey}`


        })
        .catch(error => {
            // Handle errors, e.g., display an error message
            showToast(error.message, "Error", 0);
        });



});

function CreateEmailandPassword() {
    const name = document.getElementById('admin-name').value
    const email = document.getElementById('admin-email')
    if (name == "") {
        email.value = ""
        return
    }
    email.value = (name.toLowerCase()).replace(/\s/g, '') + '@anon.com'
}

function search() {
    const searchInput = document.getElementById('Search').value.toLowerCase();
    const doctorRows = document.querySelectorAll(`.candidates-list`);

    doctorRows.forEach(row => {
        const Name = row.querySelector(`.candidate-list-title h5 a`).innerText.toLowerCase();
        if (Name.includes(searchInput)) {
            row.style.display = '';
        } else {
            row.style.display = 'none';
        }
    });
}

function DisplayFeedBacks() {
    console.log("Displaylist")
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
    document.getElementById("notapprovedseller").style.display = 'none';
    document.querySelector('.outer-container').style.display = 'none';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'block';
    document.getElementById('Inventorysnip').style.display = 'none';
    document.querySelector('.wrapper').style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.getElementById('admin-wrapper').style.display = 'none';
    document.getElementById('update-form-admin-container').style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    fetch("http://localhost:8080/getfeedback", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(),
    })
        .then(response => response.json())
        .then(data => {
            let html = ""
            console.log(data.result)
            data.result.forEach(feedback => {

                html += `

            <tr class="candidates-list">
            <td class="title">
              <div class="thumb"> <img class="img-fluid"
                  src="http://previews.123rf.com/images/jenjawin/jenjawin1904/jenjawin190400251/120265520-account-icon-outline-vector-eps10-user-profile-sign-web-icon-with-check-mark-glyph-user-authorized.jpg" alt="">
              </div>
              <div class="candidate-list-details">
                <div class="candidate-list-info">
                  <div class="candidate-list-title customer">
                    <h5 class="mb-0"><a href="#">${feedback.email}</a></h5>
                  </div>
                  <div class="candidate-list-option">
                    <ul class="list-unstyled">
                      <li><i class="fas fa-filter pr-1"></i>${feedback.role.toUpperCase()}</li>
                    </ul>
                  </div>
                </div>
              </div>
            </td>
            <td class="candidate-list-favourite-time text-center"> <a
                class="candidate-list-favourite order-2 text-danger" href="#"></a>
              <span class="candidate-list-time order-1">${feedback.feedback}</span></td>
            <td>
              <ul class="list-unstyled mb-0 d-flex justify-content-end">

     

                <li  onclick="deleteFeedback('${feedback.email}','${feedback.feedback}');DisplayFeedBacks()"><a class="text-danger" data-toggle="tooltip" title=""
                    data-original-title="Delete"><i class="far fa-trash-alt"></i></a></li>
              </ul>
            </td>
          </tr>`;

            });
            document.querySelector('.feedback-list-body').innerHTML = html;
        })
        .catch(error => {
            showToast(error, "Error", 0);
        });
}


function DisplayAllWorkers() {
    console.log("Displaylist")
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById("notapprovedseller").style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.querySelector('.outer-container').style.display = 'none';
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.getElementById('Inventorysnip').style.display = 'none';
    document.querySelector('.wrapper').style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById('update-form-admin-container').style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.getElementById('workersnip').style.display = 'block';
    fetch("http://localhost:8080/getworkers", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
    })
        .then(response => response.json())
        .then(data => {
            let html = ""
            console.log(data.result)
            data.result.forEach(worker => {

                html += `

            <tr class="candidates-list">
            <td class="title">
              <div class="thumb"> <img class="img-fluid"
                  src="data:image/jpeg;base64,${worker.image}" alt="">
              </div>
              <div class="candidate-list-details">
                <div class="candidate-list-info">
                  <div class="candidate-list-title customer">
                    <h5 class="mb-0"><a href="#">${worker.username.toUpperCase()}</a></h5>
                  </div>
                  <div class="candidate-list-option">
                    <ul class="list-unstyled">
                      <li><i class="fas fa-filter pr-1"></i>${worker.email}</li>
                    </ul>
                  </div>
                </div>
              </div>
            </td>
            <td class="candidate-list-favourite-time text-center"> <a
                class="candidate-list-favourite order-2 text-danger" href="#"></a>
              <span class="candidate-list-time order-1">${worker.no}</span></td>
            <td>
              <ul class="list-unstyled mb-0 d-flex justify-content-end">
               <li  onclick="ViewData('${worker.email}','worker');recentPage = 'worker';"><a class="text-danger" data-toggle="tooltip" title=""
              data-original-title="Delete"><i class="far fa-eye"></i></a></li>
                <li  onclick="deleteWorker('${worker.email}');DisplayAllWorkers();recentPage = 'worker';"><a class="text-danger" data-toggle="tooltip" title=""
                data-original-title="Delete"><i class="far fa-trash-alt"></i></a></li>
              </ul>
            </td>
          </tr>`;

            });
            document.querySelector('.worker-list-body').innerHTML = html;
        })
        .catch(error => {
            showToast(error, "Error", 0);
        });
}


function ViewData(id, profession) {
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.querySelector('.outer-container').style.display = 'none';
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.getElementById("notapprovedseller").style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.getElementById('Inventorysnip').style.display = 'none';
    document.querySelector('.wrapper').style.display = 'none';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'none';
    document.getElementById('update-form-admin-container').style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';

    fetch("http://localhost:8080/getdata", {
        method: "POST", // Use DELETE method to delete data
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ id: id, collection: profession })
    })

        .then(response => response.json())
        .then(data => {
            console.log(data.message)

            let html = ""
            if (profession == 'inventory') {
                html = `
                <div class="container" style="width:1500px;max-width:1500px; margin-left:300px">
                <i class="fas fa-arrow-left back-icon" onclick="BackButton()"></i>
                <div class="row">
                <div class="col-sm-8 col-sm-offset-2">
                <div class="panel panel-white profile-widget">
                <div class="row">
                <div class="col-sm-12">
                <div class="image-container bg2">
                <img src="data:image/jpeg;base64,${data.message.image}" class="avatar" alt="avatar" height="100px" >
                </div>
                </div>
                <div class="col-sm-12">
                <div class="details">
                <h4>${data.message.itemname} <i class="fa fa-sheild"></i></h4>
                <div class="mg-top-10">
                </div>
                </div>
                </div>
                </div>
                </div>
                <div class="row">
                <div class="col-sm-6"   >
                <div class="panel panel-white border-top-purple">
                <div class="panel-heading">
                <h3 class="panel-title" style="height:30px;">Item Info</h3>
                </div>
                <div class="panel-body" style="padding:30px; border-radius:5px">
                <div class="body-section">
                <h5 class="section-heading">Item Name : <span class="message">${data.message.itemname}</span></h5>
                </div>
                <div class="body-section">
                <h5 class="section-heading">Category :
                <span class="message">${data.message.itemcategory}</span>
               </h5>
               </div>
                <div class="body-section">
                <h5 class="section-heading">Price : <span class="message" >${data.message.price}  </span> </h5>
                </div>
                <div class="body-section">
                <!-- <a href="#" class="btn btn-purple btn-sm">Edit</a> -->
                </div>
                </div>
                </div>
    
                <div class="panel">
              
    
                </div>
                </div>
                <div class="col-sm-6">
                <div class="panel panel-white border-top-green">
                <div class="panel-heading">
                <h3 class="panel-title">Seller Info</h3>
              
                
                
                </div>
                <div class="panel-body" style="padding:30px">
    
    
                <div class="body-section">
                <h5 class="section-heading">Seller Name : <span class="message">  ${data.message.sellername}</span></h5>
                </div>

                <div class="body-section">
                <h5 class="section-heading">Quantity : <span class="message">${data.message.quantity}</span></h5>
                </div>
    
                <div class="body-section">
                <h5 class="section-heading">Available Quantity:  <span class="message">${data.message.sellerquantity}</span></h5>
                </div>


            
                </div>
                </div>
                <div class="panel ">
    
                </div>
      
    
                
                </div>
                </div>
                </div>
                </div>
                </div>
                </div>
            `;

            } else if (profession == 'customer') {
                html = `
                <div class="container" style="width:1500px;max-width:1500px; margin-left:300px">
                <i class="fas fa-arrow-left back-icon" onclick="BackButton()"></i>
                <div class="row">
                <div class="col-sm-8 col-sm-offset-2">
                <div class="panel panel-white profile-widget">
                <div class="row">
                <div class="col-sm-12">
                <div class="image-container bg2">
                <img src="http://cdn-icons-png.flaticon.com/512/149/149071.png" class="avatar" alt="avatar" height="100px" >
                </div>
                </div>
                <div class="col-sm-12">
                <div class="details">
                <h4>${data.message.name} <i class="fa fa-sheild"></i></h4>
                <div class="mg-top-10">
                </div>
                </div>
                </div>
                </div>
                </div>
                <div class="row">
                <div class="col-sm-6"   >
                <div class="panel panel-white border-top-purple">
                <div class="panel-heading">
                <h3 class="panel-title" style="height:30px;">Account Info</h3>
                </div>
                <div class="panel-body" style="padding:30px; border-radius:5px">
                <div class="body-section">
                <h5 class="section-heading">Account Name : <span class="message">${data.message.name}</span></h5>
                </div>
                <div class="body-section">
                <h5 class="section-heading">Password :
                <span class="message" style="display:none" id="passwordText">${data.message.password}</span>
                <a class="text-info" data-toggle="tooltip" title="Toggle Password" onclick="togglePasswordVisibility()">
                    <i class="far fa-eye" id="eyeIcon"></i>
                </a>
               </h5>
               </div>
                <div class="body-section">
                <h5 class="section-heading">ID : <span class="message" >${data.message.customerid}  </span> </h5>
                </div>
                <div class="body-section">
                <h5 class="section-heading">EmailVerified : <span class="message" >${data.message.isemailverified}  </span> </h5>
                </div>
                <div class="body-section">
                <h5 class="section-heading">Blocked : <span class="message" >${data.message.blockeduser}  </span> </h5>
                </div>
                <div class="body-section">
                <h5 class="section-heading">Wrong Attempts : <span class="message" >${data.message.wronginput} (today)</span> </h5>
                </div>
                <div class="body-section">
                <!-- <a href="#" class="btn btn-purple btn-sm">Edit</a> -->
                </div>
                </div>
                </div>
    
                <div class="panel">
              
    
                </div>
                </div>
                <div class="col-sm-6">
                <div class="panel panel-white border-top-green">
                <div class="panel-heading">
                <h3 class="panel-title">User Info</h3>
               
                </div>
                <div class="panel-body" style="padding:30px">
    
    
                <div class="body-section">
                <h5 class="section-heading">Name : <span class="message">${data.message.name}</span></h5>
                </div>
    
                <div class="body-section">
                <h5 class="section-heading">Telephone:  <span class="message">${data.message.phonenumber}</span></h5>
                </div>
                <div class="body-section">
                <h5 class="section-heading">Email : <span class="message">${data.message.email}</span></h5>
                </div>
                <div class="body-section">
                <h5 class="section-heading">Address : <span class="message">${data.message.address}</span></h5>
                </div>
            
                </div>
                </div>
                <div class="panel ">
    
                </div>
      
    
                
                </div>
                </div>
                </div>
                </div>
                </div>
                </div>
            `;
            } else if (profession == 'seller') {
                html = `
            
            <div class="container" style="width:1500px;max-width:1500px; margin-left:300px">
            <i class="fas fa-arrow-left back-icon" onclick="BackButton()"></i>
            <div class="row">
            <div class="col-sm-8 col-sm-offset-2">
            <div class="panel panel-white profile-widget">
            <div class="row">
            <div class="col-sm-12">
            <div class="image-container bg2">
            <img src="data:image/jpeg;base64,${data.message.image}" class="avatar" alt="avatar" height="100px" >
            </div>
            </div>
            <div class="col-sm-12">
            <div class="details">
            <h4>${data.message.sellername} <i class="fa fa-sheild"></i></h4>
            <div class="mg-top-10">
            </div>
            </div>
            </div>
            </div>
            </div>
            <div class="row">
            <div class="col-sm-6"   >
            <div class="panel panel-white border-top-purple">
            <div class="panel-heading">
            <h3 class="panel-title" style="height:30px;">Account Details</h3>
            </div>
            <div class="panel-body" style="padding:30px; border-radius:5px">
            <div class="body-section">
            <h5 class="section-heading">Account Name : <span class="message">${data.message.sellername}</span></h5>
            </div>
            <div class="body-section">
            <h5 class="section-heading">Password :
            <span class="message" style="display:none" id="passwordText">${data.message.password}</span>
            <a class="text-info" data-toggle="tooltip" title="Toggle Password" onclick="togglePasswordVisibility()">
                <i class="far fa-eye" id="eyeIcon"></i>
            </a>
           </h5>
           </div>
            <div class="body-section">
            <h5 class="section-heading">ID : <span class="message" >${data.message.sellerid}  </span> </h5>
            </div>
            <div class="body-section">
            <h5 class="section-heading">Blocked : <span class="message" >${data.message.blockeduser}  </span> </h5>
            </div>
            <div class="body-section">
            <h5 class="section-heading">Wrong Attempts : <span class="message" >${data.message.wronginput} (today)</span> </h5>
            </div>
            <div class="body-section">
            <!-- <a href="#" class="btn btn-purple btn-sm">Edit</a> -->
            </div>
            </div>
            </div>

            <div class="panel">
          

            </div>
            </div>
            <div class="col-sm-6">
            <div class="panel panel-white border-top-green">
            <div class="panel-heading">
            <h3 class="panel-title">Seller Info</h3>
            <div class="controls pull-right">
            <span class="pull-right clickable">
            <i class="fa fa-chevron-up"></i>
            </span>
            </div>
            </div>
            <div class="panel-body" style="padding:30px">


            <div class="body-section">
            <h5 class="section-heading">Name : <span class="message">${data.message.sellername}</span></h5>
            </div>

            <div class="body-section">
            <h5 class="section-heading">Telephone:  <span class="message">${data.message.phonenumber}</span></h5>
            </div>
            <div class="body-section">
            <h5 class="section-heading">Email : <span class="message">${data.message.selleremail}</span></h5>
            </div>
            <div class="body-section">
            <h5 class="section-heading">Address : <span class="message">${data.message.address}</span></h5>
            </div>

        
            </div>
            </div>
            <div class="panel ">

            </div>
  

            
            </div>
            </div>
            </div>
            </div>
            </div>
            </div>
            `;
            } else if (profession == 'worker') {
                html = `
                <div class="container" style="width:1500px;max-width:1500px; margin-left:300px">
               
                <div class="row">
                <i class="fas fa-arrow-left back-icon" onclick="BackButton()"></i>
                <div class="col-sm-8 col-sm-offset-2">
                <div class="panel panel-white profile-widget">
                <div class="row">
                <div class="col-sm-12">
                <div class="image-container bg2">
                <img src="data:image/jpeg;base64,${data.message.image}" class="avatar" alt="avatar" height="100px" >
                </div>
                </div>
                <div class="col-sm-12">
                <div class="details">
                <h4>${data.message.username} <i class="fa fa-sheild"></i></h4>
                <div class="mg-top-10">
                </div>
                </div>
                </div>
                </div>
                </div>
                <div class="row">
                <div class="col-sm-6"   >
                <div class="panel panel-white border-top-purple">
                <div class="panel-heading">
                <h3 class="panel-title" style="height:30px;">Account Details</h3>
                </div>
                <div class="panel-body" style="padding:30px; border-radius:5px">
                <div class="body-section">
                <h5 class="section-heading">Account Name : <span class="message">${data.message.username}</span></h5>
                </div>
                <div class="body-section">
                <h5 class="section-heading">Status :
                <span class="message"  id="passwordText">${data.message.status}</span>
               </h5>
               </div>
                <div class="body-section">
                <h5 class="section-heading">Role : <span class="message" >${data.message.role}  </span> </h5>
                </div>
                <div class="body-section">
                <!-- <a href="#" class="btn btn-purple btn-sm">Edit</a> -->
                </div>
                </div>
                </div>
    
                <div class="panel">
              
    
                </div>
                </div>
                <div class="col-sm-6">
                <div class="panel panel-white border-top-green">
                <div class="panel-heading">
                <h3 class="panel-title">Seller Info</h3>
                </div>
                <div class="panel-body" style="padding:30px">
    
    
                <div class="body-section">
                <h5 class="section-heading">Name : <span class="message">${data.message.username}</span></h5>
                </div>
    
                <div class="body-section">
                <h5 class="section-heading">Telephone:  <span class="message">${data.message.no}</span></h5>
                </div>
                <div class="body-section">
                <h5 class="section-heading">Email : <span class="message">${data.message.email}</span></h5>
                </div>
                <div class="body-section">
                <h5 class="section-heading">Salary : <span class="message">${data.message.salary}</span></h5>
                </div>
            
                </div>
                </div>
                <div class="panel ">
    
                </div>
      
    
                
                </div>
                </div>
                </div>
                </div>
                </div>
                </div>
            `;
            }


            document.querySelector('.display-view').innerHTML = html;
            document.querySelector('.display-view').style.display = 'block';
        })
        .catch(error => {
            console.log(error)
        });

}




function togglePasswordVisibility() {
    var passwordText = document.getElementById('passwordText');
    var eyeIcon = document.getElementById('eyeIcon');

    // Toggle password visibility
    if (passwordText.style.display === 'none') {
        passwordText.style.display = 'inline';
        eyeIcon.classList.remove('fa-eye-slash');
        eyeIcon.classList.add('fa-eye');
    } else {
        passwordText.style.display = 'none';
        eyeIcon.classList.remove('fa-eye');
        eyeIcon.classList.add('fa-eye-slash');
    }
}
var recentPage = ''
function BackButton() {
    if (recentPage == 'inventory') {
        DisplayListInventory()
    } else if (recentPage == 'seller') {
        DisplayListSeller()
    } else if (recentPage == 'customer') {
        DisplayListUsers()
    } else if (recentPage == 'worker') {
        DisplayAllWorkers()
    }
}

$(document).ready(async function () {
    // Initialize FullCalendar
    $('#calendar').fullCalendar({
        header: {
            left: 'prev,next',
            center: 'title',
            right: 'month,agendaWeek,agendaDay'
        },
        editable: true,
        events: await fetchEventsFromBackend(),

        eventClick: function (event) {
            showTodoList(event);
        },
        dayClick: function () {
            // Handle day click (add to-do item)
            DisplayEventForm();
        }

    });

    async function fetchEventsFromBackend() {
        var adminData = localStorage.getItem('admindata');
        var adminObject = JSON.parse(adminData);
        const data = {
            email: adminObject.username,
        };

        try {
            const response = await fetch("http://localhost:8080/getevent", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data),
            });

            const apiResponse = await response.json();
            const fullCalendarEvents = convertApiResponseToFullCalendarEvents(apiResponse.message);

            console.log('FullCalendar Events:', fullCalendarEvents); // Log the events before returning

            return fullCalendarEvents;
        } catch (error) {
            return [];
        }
    }

    function convertApiResponseToFullCalendarEvents(apiResponse) {
        console.log(apiResponse)
        return apiResponse.map(event => {
            console.log('Converted Event:', event.title);
            return {
                title: event.title,
                start: event.start,
                end: event.end,
                todos: event.todos
            };
        });
    }

    function showTodoList(event) {
        const todos = event.todos || [];
        let todoList = `To-Do :`;
        if (todos.length === 0) {
            todoList += 'No items.';
        } else {
            todos.forEach((item, index) => {
                todoList += ` ${item}`;
            });
        }
        showToast(todoList, "Info", 3)
    }



});
function DisplayCalender() {
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    document.querySelector('.wrapper').style.display = 'none';
    document.querySelector('.outer-container').style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById("notapprovedseller").style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById('Inventorysnip').style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'none';
    document.getElementById("calendar").style.display = 'block';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.getElementById('update-form-admin-container').style.display = 'none';
}

function DisplayEventForm() {
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    document.querySelector('.wrapper').style.display = 'none';
    document.querySelector('.outer-container').style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.getElementById("notapprovedseller").style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.getElementById('Inventorysnip').style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'none';
    document.getElementById("calendar").style.display = 'block';
    document.getElementById('update-form-admin-container').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'block';
}

document.getElementById("event-form").addEventListener("submit", function (event) {
    event.preventDefault();

    const start = document.getElementById("event-start").value;
    const end = document.getElementById("event-end").value;
    const title = document.getElementById("event-title").value;
    const todo = document.getElementById("event-todo").value;
    var adminData = localStorage.getItem('admindata');
    var adminObject = JSON.parse(adminData);
    console.log(adminData.username)

    const requestData = {
        email: adminObject.username,
        start,
        end,
        title,
        todos: [todo]
    };
    console.log(requestData)

    // Send a DELETE request to your server to delete the data
    fetch("http://localhost:8080/addevent", {
        method: "POST", // Use DELETE method to delete data
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(requestData)
    })
        .then(response => response.json())
        .then(data => {

            if (data.message) {
                showToast(data.message, "Success", 3)
                document.getElementById("event-form").reset()
            } else if (data.error) {
                showToast(data.error, "Warning", 0)
            }
        })
        .catch(error => {
            showToast(data.error, "Warning", 0)
        });

});

function BlockUser(email, collection) {
    const formData = {
        email,
        collection,
    }
    fetch("http://localhost:8080/block", {
        method: "POST", // Use DELETE method to delete data
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
    })
        .then(response => response.json())
        .then(data => {

            if (data.message) {
                showToast(data.message, "Success", 3)
            } else if (data.error) {
                showToast(data.error, "Warning", 0)
            }
        })
        .catch(error => {
            showToast(data.error, "Warning", 0)
        });
}

function DisplayShutDown() {
    document.getElementById("shutdown-form-container").style.display = 'block';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.querySelector('.wrapper').style.display = 'none';
    document.querySelector('.outer-container').style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById('Inventorysnip').style.display = 'none';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.getElementById("notapprovedseller").style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'none';
    document.getElementById("calendar").style.display = 'block';
    document.getElementById('update-form-admin-container').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.getElementById("clearbuttons-container").style.display = 'none';
}

document.getElementById("shutdown-form").addEventListener("submit", () => {
    var adminData = localStorage.getItem('admindata');
    var adminObject = JSON.parse(adminData);
    const formData = {
        token: adminObject.token,
        password: (document.getElementById("secretkey").value).trim()
    }
    if (formData.password == "") {
        showToast("Key Required to ShutDown", "Danger", 0)
        return
    }

    fetch("http://localhost:8080/shutdown", {
        method: "POST", // Use DELETE method to delete data
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
    })
        .then(response => response.json())
        .then(data => {

            if (data.message) {
                showToast(data.message, "Success", 3)
            } else if (data.error) {
                showToast(data.error, "Warning", 0)
            }
        })
        .catch(error => {
            showToast(data.error, "Warning", 0)
        });
})

function DisplayClearForm() {
    document.getElementById("shutdown-form-container").style.display = 'none'
    document.getElementById('workersnip').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.querySelector('.wrapper').style.display = 'none';
    document.querySelector('.outer-container').style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById('Inventorysnip').style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
    document.getElementById('update-form-admin-container').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.getElementById("notapprovedseller").style.display = 'none';
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'block'
}

document.getElementById("cleardb-form").addEventListener("submit", () => {

    const password = (document.getElementById("dbsecretkey").value).trim()
    if (password == "AnonClearDB") {
        DisplayClearData()
    } else {
        showToast("Invalid Key", "Error", 3)
    }
})

function DisplayClearData() {
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.querySelector('.wrapper').style.display = 'none';
    document.querySelector('.outer-container').style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById('Inventorysnip').style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById('update-form-admin-container').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.getElementById("orders-container").style.display = 'none';
    document.getElementById("notapprovedseller").style.display = 'none';
    document.getElementById("clearbuttons-container").style.display = 'block';
}

function ClearDB(collection) {
    var adminData = localStorage.getItem('admindata');
    var adminObject = JSON.parse(adminData);
    const formData = {
        id: adminObject.token,
        collection,
    }
    fetch("http://localhost:8080/cleardb", {
        method: "POST", // Use DELETE method to delete data
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
    })
        .then(response => response.json())
        .then(data => {

            if (data.message) {
                showToast(data.message, "Success", 3)
            } else if (data.error) {
                showToast(data.error, "Warning", 0)
            }
        })
        .catch(error => {
            showToast(data.error, "Warning", 0)
        });
}

// Display Conformation

function showConfirmation(function_name, question, option1, option2, id) {
    console.log("In Conformation")
    document.getElementById("conformationoverlay").classList.add("conformationactive");
    document.getElementById("confirmationDialog").classList.add("conformationactive");
    document.querySelector(".conformation-question").innerHTML = question;
    document.getElementById("conformationtrue").innerHTML = option1;
    document.getElementById("confirmationfalse").innerHTML = option2;

    document.getElementById("conformationtrue").addEventListener("click", function () {
        function_name(id);
        hideConfirmationDialog();
    });
    document.getElementById("confirmationfalse").addEventListener("click", function () {
        hideConfirmationDialog();
    });
}

function hideConfirmationDialog() {
    document.getElementById("conformationoverlay").classList.remove("conformationactive");
    document.getElementById("confirmationDialog").classList.remove("conformationactive");
}

function Displayallnotapprovedseller() {
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.querySelector('.wrapper').style.display = 'none';
    document.querySelector('.outer-container').style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById('Inventorysnip').style.display = 'none';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById('update-form-admin-container').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
    document.getElementById("notapprovedseller").style.display = 'block';
    var adminData = localStorage.getItem('admindata');
    var adminObject = JSON.parse(adminData);
    const formData = {
        token: adminObject.token,
    }
    fetch("http://localhost:8080/getallnotapprovedseller", {
        method: "POST", // Use DELETE method to delete data
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
    })
        .then(response => response.json())
        .then(data => {
            document.querySelector('.notapprovedseller-list-body').innerHTML = "";
            if (data.message) {
                let html = ""
                data.message.forEach((items) => {
                    html += `
                    
            <tr class="candidates-list">
            <td class="title">
              <div class="thumb"> <img class="img-fluid"
                  src="data:image/jpeg;base64,${items.image}" alt="">
              </div>
              <div class="candidate-list-details">
                <div class="candidate-list-info">
                  <div class="candidate-list-title customer">
                    <h5 class="mb-0"><a href="#">${items.sellername}</a></h5>
                  </div>
                  <div class="candidate-list-option">
                    <ul class="list-unstyled">
                      <li><i class="fas fa-filter pr-1"></i>${items.selleremail}</li>
                    </ul>
                  </div>
                </div>
              </div>
            </td>
            <td class="candidate-list-favourite-time text-center"> <a
                class="candidate-list-favourite order-2 text-danger" href="#"></a>
              <span class="candidate-list-time order-1">${items.phoneno}</span></td>
            <td>
              <ul class="list-unstyled mb-0 d-flex justify-content-end">
               <li  onclick="ViewData('${items.selleremail}','seller');recentPage = 'worker';"><a class="text-danger" data-toggle="tooltip" title=""
              data-original-title="Delete"><i class="far fa-eye"></i></a></li>
                <li  onclick="DeleteData('${items.selleremail}','seller');recentPage = 'worker';"><a class="text-danger" data-toggle="tooltip" title=""
                data-original-title="Delete"><i class="far fa-trash-alt"></i></a></li>
                <li  onclick="ApproveSeller('${items.sellerid}');recentPage = 'worker';"><a class="text-danger" data-toggle="tooltip" title=""
                data-original-title="Delete"><i class="far fa-check-circle approve-icon"></i></a></li>
              </ul>
            </td>
          </tr>
                    `

                })
                document.querySelector('.notapprovedseller-list-body').innerHTML = html
            } else if (data.error) {
                showToast(data.error, "Warning", 0)
            }
        })
        .catch(error => {
            showToast(error, "Warning", 0)
        });
}

function ApproveSeller(id) {
    var adminData = localStorage.getItem('admindata');
    var adminObject = JSON.parse(adminData);
    const formData = {
        token: adminObject.token,
        sellerid: id,
    }
    fetch("http://localhost:8080/approveseller", {
        method: "POST", // Use DELETE method to delete data
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
    })
        .then(response => response.json())
        .then(data => {

            if (data.message) {
                showToast(data.message, "Success", 3)
                Displayallnotapprovedseller();
            } else if (data.error) {
                showToast(data.error, "Warning", 0)
            }
        })
        .catch(error => {
            showToast(data.error, "Warning", 0)
        });
}

function DisplayBlock() {
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.querySelector('.wrapper').style.display = 'none';
    document.querySelector('.outer-container').style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById('Inventorysnip').style.display = 'none';
    document.getElementById("block-form-container").style.display = 'block';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.getElementById('update-form-admin-container').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.getElementById("notapprovedseller").style.display = 'none';
    document.getElementById("orders-container").style.display = 'none';
}

function BLockbyemail(event) {
    event.preventDefault();
    let email = (document.getElementById("blockemail").value).trim()
    let collection = document.getElementById("blockcollection").value;
    console.log(email, collection)
    BlockUser(email, collection)
    document.getElementById("blockemail").value = ""
}

function DisplayOrders() {
    document.getElementById("shutdown-form-container").style.display = 'none';
    document.getElementById('employee-wrapper').style.display = 'none';
    document.getElementById('workersnip').style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.querySelector('.wrapper').style.display = 'none';
    document.querySelector('.outer-container').style.display = 'none';
    document.querySelector('.container-p-y').style.display = 'none';
    document.getElementById('snippetContent').style.display = 'none';
    document.getElementById('sellersnip').style.display = 'none';
    document.getElementById('Inventorysnip').style.display = 'none';
    document.getElementById("block-form-container").style.display = 'none';
    document.getElementById('feedbacksnip').style.display = 'none';
    document.querySelector('.display-view').style.display = 'none';
    document.getElementById('admin-wrapper').style.display = 'none';
    document.getElementById('single-order-container').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById('update-form-admin-container').style.display = 'none';
    document.getElementById("calendar").style.display = 'none';
    document.getElementById("event-wrapper").style.display = 'none';
    document.getElementById("clear-form-container").style.display = 'none'
    document.getElementById("clearbuttons-container").style.display = 'none';
    document.getElementById("notapprovedseller").style.display = 'none';
    document.getElementById("orders-container").style.display = 'block';
    var adminData = localStorage.getItem('admindata');
    var adminObject = JSON.parse(adminData);
    const formData = {
        token: adminObject.token,
    }

    fetch("http://localhost:8080/getallorders", {
        method: "POST", // Use DELETE method to delete data
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
    })

        .then(response => response.json())
        .then(data => {
            let html = ""
            data.message.forEach(items => {

                html += `
            <tr class="candidates-list seller-list">
            <td class="title">
              <div class="thumb"> <img class="img-fluid"
                  src="data:image/jpeg;base64,${items.itemstobuy.image}" alt="">
              </div>
              <div class="candidate-list-details">
                <div class="candidate-list-info">
                  <div class="candidate-list-title seller">
                    <h5 class="mb-0"><a href="#">${items.itemstobuy.productname.toUpperCase()}</a></h5>
                  </div>
                  <div class="candidate-list-option">
                    <ul class="list-unstyled">
                      <li><i class="fas fa-filter pr-1"></i>${items.itemstobuy.itemcategory}
                      </li>
                      <li><i class="fas fa-map-marker-alt pr-1"></i>${items.itemstobuy.quantity}</li>
                    </ul>
                  </div>
                </div>
              </div>
            </td>
            <td class="candidate-list-favourite-time text-center"> <a
                class="candidate-list-favourite order-2 text-danger" href="#"></a>
              <span class="candidate-list-time order-1">${items.customerid}</span></td>
            <td>
              <ul class="list-unstyled mb-0 d-flex justify-content-end">
                    <li onclick="GetOrder('${items.orderid}','seller');recentPage = 'inventory';"><a  class="text-info" data-toggle="tooltip" title="" data-original-title="Edit"><i
                    class="fas fa-eye"></i></a>
                    </li>
                <li  onclick="DeleteOrder('${items.orderid}','sel');recentPage = 'inventory';"><a class="text-danger" data-toggle="tooltip" title=""
                    data-original-title="Delete"><i class="far fa-trash-alt"></i></a></li>
              </ul>
            </td>
          </tr>`;

            });
            document.querySelector('.orders-list-body').innerHTML = html;
        })
        .catch(error => {
            showToast(error, "Error", 0);
        });
}


async function GetOrder(id) {
    try {
        document.getElementById("shutdown-form-container").style.display = 'none';
        document.getElementById('employee-wrapper').style.display = 'none';
        document.getElementById('workersnip').style.display = 'none';
        document.getElementById("clear-form-container").style.display = 'none'
        document.querySelector('.wrapper').style.display = 'none';
        document.querySelector('.outer-container').style.display = 'none';
        document.querySelector('.container-p-y').style.display = 'none';
        document.getElementById('snippetContent').style.display = 'none';
        document.getElementById('sellersnip').style.display = 'none';
        document.getElementById('Inventorysnip').style.display = 'none';
        document.getElementById("block-form-container").style.display = 'none';
        document.getElementById('feedbacksnip').style.display = 'none';
        document.querySelector('.display-view').style.display = 'none';
        document.getElementById('admin-wrapper').style.display = 'none';
        document.getElementById('single-order-container').style.display = 'none';
        document.getElementById("calendar").style.display = 'none';
        document.getElementById('update-form-admin-container').style.display = 'none';
        document.getElementById("calendar").style.display = 'none';
        document.getElementById("event-wrapper").style.display = 'none';
        document.getElementById("clear-form-container").style.display = 'none'
        document.getElementById("clearbuttons-container").style.display = 'none';
        document.getElementById("notapprovedseller").style.display = 'none';
        document.getElementById("orders-container").style.display = 'none';
        var adminData = localStorage.getItem('admindata');
        var adminObject = JSON.parse(adminData);
        const data = {
            token: adminObject.token,
            orderid: id,
        }
        console.log(data)
        const output = await fetch('http://localhost:8080/getcustomerorderforadmin', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });
        const value = await output.json();
        console.log(value)
        if (value.message) {

            let html = ""
            let orderstatus = ""
            const item = value.message
            console.log(value)
            const entries = Object.entries(item.status);

            entries.forEach(([key, val]) => {
                if (val == "completed") {
                    orderstatus = key
                    return
                }
            });

            html = `<div class="container">
        
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
                    <!-- <span class="me-3">CASH ON DELIVERY</span> -->
                    <span class="badge rounded-pill bg-info">${orderstatus.toUpperCase()}</span>
                    <button class="${item.status.dispatched}-cancelbutton" onclick="showConfirmation(DeleteOrder,'Are you sure want to Cancel this order ?','Yes','No','${item.orderid}')">Cancel Order</button>
                  </div>
                </div>
                <table class="table table-borderless">
                  <tbody>
                    <tr>
                      <td>
                        <div class="d-flex mb-2">
                          <div class="flex-shrink-0">
                            <img src="data:image/jpeg;base64,${item.itemstobuy.image}" alt width="70px" class="img-fluid">
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
                  <hr>
                    <tr class="fw-bold">
                      
                      <td colspan="2">TOTAL</td>
                      <td class="text-end">₹${item.totalamount}</td>
                    </tr>
                  </tfoot>
                </table>
              </div>
            </div>
  
  
            <div class="card-body"  style="padding:0px">
              <div class="row" style="padding:0px">
  
                <div class="card mb-3">
                  <div class="d-flex flex-wrap flex-sm-nowrap justify-content-between py-3 px-2 bg-secondary" style="color:white;width:100%;margin-left:0px;margin-right:0px;max-width:100%">
                    <div class="w-100 text-center py-1 px-2"><span class="text-medium">Shipped By:</span>
                      ${item.itemstobuy.sellerid}
                    </div>
                    <div class="w-100 text-center py-1 px-2"><span class="text-medium">Status:</span>
                      ${orderstatus.toUpperCase()}
                    </div>
                    <div class="w-100 text-center py-1 px-2"><span class="text-medium">Expected
                        Date:</span>${item.deliverydate}
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
              <h3 class="h6">Admin Notes</h3>`
            if (item.status.dispatched == "pending") {
                html += ` <p>Seller is processing the order.</p>
                <p>Order Cancellation will affect Seller & Customer rateing</p>
                <p>Any problem feel free to contact us</p>`
            } else {
                html += ` <p>🎉Order Delivered Successfully 🎊 </p>
                <p>Custmer can return if needed🤩🤩</p>
                <p>Any problem feel free to contact us</p>`
            }

            html += `
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
            
            document.getElementById('single-order-container').innerHTML = html;
            document.getElementById('single-order-container').style.display = 'block';

        } else if (value.error) {
            showToast(value.error, "Error", 0)
        }

    } catch (error) {
        console.log(error)
    }
}

// Async function to delete order
async function DeleteOrder(id) {
    try {
        var adminData = localStorage.getItem('admindata');
        var adminObject = JSON.parse(adminData);
        const data = {
            token: adminObject.token,
            orderid: id,
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
        if (value.message) {
            showToast(value.message, "Success", 3)
            DisplayOrders()
        } else if (value.error) {
            showToast(value.error, "Error", 0)
        }
    } catch (error) {
        console.log(error)
    }
}



