<div class="container">
  <div class="logo">
    <img draggable={false} src="/soar-logo.svg" alt="SOAR Logo" />
  </div>
  <div class="header-text">
    <h1>Avionics Propulsion Pipeline</h1>
    <p>
      Automating the post-processing of propulsion test data and archiving
      results online for analysis.
    </p>
  </div>
</div>

<style scoped lang="scss">
  @use "../styles/variables.scss" as *;

  div.container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }

  div.logo img {
    height: 4rem;
    width: auto;
  }

  div.header-text {
    text-align: center;
    display: flex;
    flex-direction: column;
    align-items: center;

    h1 {
      color: $txt-color-1;
      font-size: 2.5rem;
      margin: 0.5rem 0;
    }

    p {
      color: $txt-color-2;
      max-width: 30rem;
      font-size: 1.1rem;
      margin: 0.1rem 0.5rem 0.5rem;
    }
  }
</style>
