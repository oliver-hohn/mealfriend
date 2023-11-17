'use client';

import { useState } from 'react';

import { BrowserHeaders } from "browser-headers";

import { MealfriendClient } from "../protos/mealfriend_pb_service"
import { GetMealPlanRequest, GetMealPlanResponse, Recipe } from "../protos/mealfriend_pb"
import RecipeView from './RecipeView';


export default function Home() {
  const [recipes, setRecipes] = useState<Recipe[]>([]);

  const handleRequestMealPlanClick = () => {
    const client = new MealfriendClient('http://localhost:8080');

    const request = new GetMealPlanRequest();
    request.getRequirementsMap().set('unspecified', 5);

    client.getMealPlan(request, new BrowserHeaders(), (err, response) => {
      if (err) {
        console.error(err);
        return;
      }

      if (response === null) {
        console.error('response is null');
        return;
      }

      setRecipes(response.getRecipesList());
    });
  };

  const hasRecipes = recipes.length > 0;

  const recipeViews = recipes.map((recipe, index) => {
    return (
      <RecipeView key={index} recipe={recipe} />
    );
  });

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      { hasRecipes ? (
        <div className="grid grid-cols-3 gap-x-5 gap-y-10 max-w-5xl v">
          {recipeViews}
        </div>
      ) : (
        <div className = "mb-32 grid text-center lg:max-w-5xl lg:w-full lg:mb-0 lg:grid-cols-1 lg:text-left">
          <button
            className = "bg-blue-800 hover:bg-blue-700 text-white font-bold py-4 px-4 rounded-full"
            onClick = { handleRequestMealPlanClick }
          >
            <h2 className = {`mb-3 text-2xl font-semibold`}>
              Surprise me!
            </h2>
          </button >
        </div >
      ) }

    </main>
  )
}
