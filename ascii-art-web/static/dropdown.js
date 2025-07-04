function initDropdown() {
    const toggle = document.getElementById("banner-toggle");
    const menu = document.getElementById("banner-menu");
    const hidden = document.getElementById("banner");
  
    toggle.addEventListener("click", () => {
      menu.style.display = menu.style.display === "block" ? "none" : "block";
    });
  
    menu.querySelectorAll("li").forEach(item => {
      item.addEventListener("click", () => {
        const value = item.getAttribute("data-value");
        hidden.value = value;
        toggle.textContent = item.textContent;
        menu.style.display = "none";
        scheduleGenerate();
      });
    });
  
    document.addEventListener("click", e => {
      if (!document.getElementById("banner-dropdown").contains(e.target)) {
        menu.style.display = "none";
      }
    });
  }
  