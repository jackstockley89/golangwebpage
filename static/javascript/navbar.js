var navbarTitle = document.getElementById('navbar-container');
var navbar = `
<header>
  <div class="container">
    <div class="topnav">
      <a href="/home" class="back-link">Home</a>
      <a href="/activities" class="back-link">Activities</a>
    </div>
  </div>
</header>`;
document.body.insertAdjacentHTML('afterbegin', navbar);