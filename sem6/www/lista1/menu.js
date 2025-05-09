document.addEventListener("DOMContentLoaded", () => {
  const header = document.querySelector("header");
  const nav = header.querySelector("nav");

  // 1) Tworzymy przycisk „hamburger”
  const btn = document.createElement("button");
  btn.className = "hamburger";
  btn.setAttribute("aria-label", "Toggle menu");
  btn.innerHTML = "&#9776;"; // ☰
  header.insertBefore(btn, nav);

  // 2) Toggle klasy .open przy kliknięciu
  btn.addEventListener("click", () => {
    nav.classList.toggle("open");
  });

  // 3) Przy resize: decydujemy, czy pokazać nav czy hamburger
  const updateMenu = () => {
    if (window.innerWidth <= 600) {
      btn.style.display = "block";
      if (!nav.classList.contains("open")) nav.style.display = "none";
    } else {
      btn.style.display = "none";
      nav.style.display = "flex";
      nav.classList.remove("open");
    }
  };

  window.addEventListener("resize", updateMenu);
  updateMenu();
});
