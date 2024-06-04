<!DOCTYPE html>
<html>
  <head>
    <script>
      function fact() {
        var i, num, f;
        f = 1;
        num = document.getElementById("num").value;
        for (i = 1; i <= num; i++) {
          f = f * i;
        }
        document.getElementById("res").innerHTML = "The factorial of the number " + num + " is: " + f;
      }
    </script>
  </head>
  <body style="text-align: center; font-size: 20px;">
    <h1>Welcome to the javaScript world!!</h1>
    Enter a particular number: <input id="num"><br><br>
    <button onclick="fact()">Please type any Factorial number</button>
    <p id="res"></p>
  </body>
</html>
