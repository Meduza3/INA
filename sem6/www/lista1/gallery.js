const imageUrls = ["warhammer.png", "minecraft.png"];

function loadImage(url) {
  return new Promise((resolve, reject) => {
    const img = new Image();
    img.src = url;
    img.onload = () => resolve(img);
    img.onerror = () => reject(new Error(`Nie udało się załadować ${url}`));
  });
}

// Równoległe ładowanie wszystkich zdjęć
Promise.all(imageUrls.map(loadImage))
  .then((images) => {
    const gallery = document.getElementById("gallery");
    images.forEach((img) => {
      const wrapper = document.createElement("div");
      wrapper.classList.add("gallery-item");
      wrapper.appendChild(img);
      gallery.appendChild(wrapper);
    });
  })
  .catch((error) => console.error(error));
