function fetchAndDisplayRecipes() {
    fetch("/recipes")
      .then((response) => response.json())
      .then((recipes) => {
        const container = document.getElementById("recipes-container");
        container.innerHTML = ""; // コンテナをクリア
        recipes.forEach((recipe) => {
          const recipeElement = document.createElement("div");
          recipeElement.innerHTML = `
          <h2>${recipe.title}</h2>
          <p>${recipe.description}</p>
          <ul>
            ${recipe.ingredients
              .map((ingredient) => `<li>${ingredient}</li>`)
              .join("")}
          </ul>
          <ol>
            ${recipe.instructions
              .map((instruction) => `<li>${instruction}</li>`)
              .join("")}
          </ol>
        `;
          container.appendChild(recipeElement);
        });
      })
      .catch((error) => console.error("Error:", error));
  }
  
  // ページの読み込みが完了したらレシピを取得
  document.addEventListener("DOMContentLoaded", fetchAndDisplayRecipes);