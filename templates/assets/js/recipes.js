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

// 非表示にする
function hideElement(element) {
  element.innerHTML = "";
}

// ページの読み込みが完了したらレシピを取得
document
  .getElementById("fetch-recipes-button")
  .addEventListener("click", fetchAndDisplayRecipes);

document.getElementById("hide-recipes-button").addEventListener("click", () => {
  const container = document.getElementById("recipes-container");
  hideElement(container);
});

// 特定のIDのレシピを取得して表示する関数
function fetchAndDisplayRecipeById() {
  const id = document.getElementById("recipe-id-input").value;
  if (id) {
    fetch(`/recipes/${id}`)
      .then((response) => {
        if (!response.ok) {
          throw new Error("レシピが見つかりませんでした");
        }
        return response.json();
      })
      .then((recipe) => {
        const container = document.getElementById("recipes-container");
        container.innerHTML = `
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
      })
      .catch((error) => {
        console.error("Error:", error);
        document.getElementById("recipes-container").innerHTML = error.message;
      });
  } else {
    document.getElementById("recipes-container").innerHTML =
      "IDを入力してください";
  }
}

// ボタンのクリックイベントリスナーを追加
document
  .getElementById("fetch-recipe-button")
  .addEventListener("click", fetchAndDisplayRecipeById);
