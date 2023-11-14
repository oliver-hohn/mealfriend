'use client';

import React from 'react';

import { BrowserHeaders } from "browser-headers";

import { MealfriendClient } from "../protos/mealfriend_pb_service"
import { GetMealPlanRequest, GetMealPlanResponse } from "../protos/mealfriend_pb"

export default function Home() {
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

      console.log(`response: ${response.getRecipesList()}}`);
    });
  };

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div className="mb-32 grid text-center lg:max-w-5xl lg:w-full lg:mb-0 lg:grid-cols-1 lg:text-left">
        <button
          className="bg-blue-800 hover:bg-blue-700 text-white font-bold py-4 px-4 rounded-full"
          onClick={handleRequestMealPlanClick}
        >
          <h2 className={`mb-3 text-2xl font-semibold`}>
            Surprise me!
          </h2>
        </button>
      </div>
    </main>
  )
}
