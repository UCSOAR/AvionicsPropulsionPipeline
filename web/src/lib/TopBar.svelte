<header class="top-bar">
  <div class="logo">
    <img src="/soar-logo.svg" alt="" />
  </div>

  <nav class="nav-links">
    <a href="#">pyTac</a>
    <a href="#">pyHac</a>
    <a href="#">mapleleaf</a>
  </nav>
</header>

<style scoped lang="scss">
  @use "../styles/variables.scss" as *;

  .top-bar {
    display: flex;
    align-items: center;
    padding: 0 1rem;
    height: 3.75rem;
    background-color: black;
  }

  .logo {
    margin-right: auto; /* Moves logo to the left */
  }

  .nav-links {
    display: flex;
    justify-content: center; /* Centers links inside */
    flex-grow: 1; /* Allows nav-links to take available space */
    gap: 2rem;
  }

  .nav-links a {
    text-decoration: none;
    color: white; /* Light text */
    font-size: 1.2rem;
    font-weight: 500;
    transition: color 0.2s ease;
  }

  .logo img {
    height: 3rem; /* Restores original size */
    width: auto; /* Keeps aspect ratio */
  }

  .nav-links a:hover {
    color: red; /* Change color on hover */
  }
</style>
