import { Recipe } from "../protos/mealfriend_pb"

import NewTabIcon from '../assets/new_tab.svg';

import * as google_protobuf_duration_pb from "google-protobuf/google/protobuf/duration_pb";

interface RecipeViewProps {
  recipe: Recipe;
}

function formatDuration(duration: (google_protobuf_duration_pb.Duration | undefined)): string {
  if (!duration) {
    return "Unknown";
  }

  const minutes = Math.floor(duration.getSeconds() / 60);
  const hours = Math.floor(minutes / 60);

  if (hours === 0) {
    return `${minutes}mins`;
  }

  let output = "";
  if (hours === 1) {
    output += `${hours}hr`;
  } else {
    output += `${hours}hrs`;
  }

  const remainingMins = minutes % 60;

  if (remainingMins > 0) {
    output += ` ${remainingMins}mins`;
  }

  return output;
}

function formatSource(source: string): string {
  const url = new URL(source);
  let hostname = url.hostname;
  hostname = hostname.replace(/^www\./, '')

  return hostname;
}

export default function RecipeView({ recipe }: RecipeViewProps) {
  const handleViewSourceClick = () => {
    window.open(recipe.getSource(), '_blank');
  };

  const handleRefreshClick = () => {
    console.log("refresh clicked");
  };

  return (
    <div className="grid grid-cols-3 gap-y-3">
      <div className="col-span-full">
        <img className="rounded-2xl w-full h-48 object-cover" src={recipe.getImageUrl()} alt={recipe.getName()} />
      </div>
      <div className="col-span-full flex items-center justify-center">
        <span className="font-mono text-lg text-center">{recipe.getName()}</span>
      </div>
      <div className="col-span-full flex items-center justify-center">
        <span className="font-mono text-sm text-center">Cook time: {formatDuration(recipe.getCookTime())}</span>
      </div>
      <div className="col-span-full">
        <div className="flex items-center justify-center">
          <button className="source-button rounded-full w-1/2 h-full flex items-center justify-center p-2" onClick={handleViewSourceClick}>
            <span>{formatSource(recipe.getSource())}</span>
            <NewTabIcon className="fill-white h-5 w-5 ml-2" />
          </button>
        </div>
      </div>
    </div>
  );
}
