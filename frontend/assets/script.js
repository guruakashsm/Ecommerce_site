const list = document.querySelectorAll(".list");
function activeLink() {
  list.forEach((item) => item.classList.remove("active"));
  this.classList.add("active");
}
list.forEach((item) => item.addEventListener("click", activeLink));

// Create a mapping of text to icon names
const iconMapping = {
  Home: "home",
  Signin: "person",
  Signup: "person-add",
  Admin: "shield-checkmark",
  Seller: "briefcase",
  Feedback: "chatbubble-ellipses",
};

// Update icons based on text
const icons = document.querySelectorAll(".icon ion-icon");
icons.forEach((icon) => {
  const text = icon.closest(".list").querySelector(".text").textContent;
  if (iconMapping.hasOwnProperty(text)) {
    icon.setAttribute("name", iconMapping[text]);
  }
});
