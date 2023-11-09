const { GetMealPlanRequest } = require('./protos/mealfriend_pb.js');
const { MealfriendClient } = require('./protos/mealfriend_grpc_web_pb.js');

var client = new MealfriendClient('http://localhost:8080', null, null);

var request = new GetMealPlanRequest();
request.getRequirementsMap().set("beef", 1).set("poultry", 1)

client.getMealPlan(request, {}, (err, response) => {
    if (err) {
        console.log(`Unexpected error for scrape: code = ${err.code}` +
            `, message = "${err.message}"`);
    } else {
        console.log(response.getRecipesList());
    }
});
