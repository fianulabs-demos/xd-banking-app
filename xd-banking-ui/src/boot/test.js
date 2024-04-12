function calculateExpression(input) {
  try {
    let result = eval(input);
    console.log("Result:", result);
  } catch (error) {
    console.error("Error evaluating expression:", error);
  }
}

const userInput = "3 + 2"; // Replace this string with any expression
calculateExpression(userInput);
