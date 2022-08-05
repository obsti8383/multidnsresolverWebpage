<!DOCTYPE html>

<html>
<head>
  <title>Multi DNS Resolver</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }

    header,
    footer {
      width: 960px;
      margin-left: auto;
      margin-right: auto;
    }

    header {
      padding: 100px 0;
    }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 50px 0;
      color: #999;
    }

    .description {
      text-align: center;
      font-size: 16px;
    }

    table, th, td {
      border: 1px solid white;
      border-collapse: collapse;
      font-size: 16px;
      border-radius: 10px;
      margin-left: auto;
      margin-right: auto;
    }

    a {
      color: #444;
      text-decoration: none;
    }

    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }

    .headline{
      text-align: center;
      font-size: 36px;
    }
  </style>
</head>

<body>
  <header>  
    <h1 class="headline">Multi DNS Resolver</h1>
  </header>

  <div class="description">
    <form action="/" method="post">
      Servername: <input type="text" name="Servername" autofocus value="{{.Servername}}"></td>
      <input type="submit" value="Go">
    </form>
    <br><br>
    <textarea cols="150" rows="6" readonly>{{.Output}}</textarea>
  </div>
  

<!-- <form action="/" method="post">
  <table>
  <tr>
  <td>Start: <input type="time" name="start" value="{{.Start}}"></td>
  <td>End: <input type="time" name="end" value="{{.End}}"></td>
  <td>TestInput: <input type="text" name="TestInput" value="{{.TestInput}}"></td>
  </tr>
  <tr>
  <td>Start: <input type="time" name="start2" value="{{.Start}}"></td>
  <td>End: <input type="time" name="end2" value="{{.End}}"></td>
  <td>TestInput: <input type="text" name="TestInput2" value="{{.TestInput2}}"></td>
  </tr>
  </table>
  <div class="description"><input type="submit" value="Save"></div>
</form> -->
  <footer>
    <div class="author">
      Official website:
      <a href="http://{{.Website}}">{{.Website}}</a> /
      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
  </footer>
  <div class="backdrop"></div>

<!-- <script src="/static/js/reload.min.js"></script> --->
</body>
</html>
