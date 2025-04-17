<script lang="ts">
  import "$lib/styles/global.scss";
  import { goto } from "$app/navigation";
  import { LucideX, Rocket } from "@lucide/svelte";
  import IconButton from "$lib/components/IconButton.svelte";

  function goToTAC() {
    goto("/tac");
  }
</script>

<main>
  <header class="front-page">
    <div class="logo">
      <img draggable={false} src="/soar-logo.svg" alt="SOAR Logo" />
    </div>
  </header>

  <div class="content-container">
    <div class="header-text">
      <h1>Avionics Propulsion Pipeline</h1>
      <p>
        Automating the post-processing of propulsion test data and<br />
        archiving results online for analysis.
      </p>
    </div>

    <div class="card-container">
      <div class="card">
        <h2>TAC</h2>
        <p>
          <b>Test Analysis Code</b> developed for the purpose of analysing data obtained
          from hybrid engine tests.
        </p>
        <div class="button-container">
          <IconButton icon={Rocket} label="Get Started" onClick={goToTAC} />
        </div>
      </div>

      <div class="card">
        <h2>HAC</h2>
        <p><b>Hybrid Analysis Code</b>...</p>
        <div class="button-container">
          <IconButton
            icon={LucideX}
            label="Not Available"
            isDisabled={true}
            onClick={() => {}}
          />
        </div>
      </div>

      <div class="card">
        <h2>Mapleleaf</h2>
        <p>Simulation software used for rocket testing.</p>
        <div class="button-container">
          <IconButton
            icon={LucideX}
            label="Not Available"
            isDisabled={true}
            onClick={() => {}}
          />
        </div>
      </div>
    </div>
  </div>
</main>

<style lang="scss">
  @use "../lib/styles/variables.scss" as *;

  main {
    background-color: $bg-color-1;
    height: 100%;
  }

  header {
    width: 100%;
    display: flex;
    justify-content: center;
    padding-top: 1rem;
  }

  div.content-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 4rem 0;
    gap: 1rem;
    height: 100%;
  }

  .logo img {
    height: 4rem;
    width: auto;
  }

  .header-text {
    color: $txt-color-1;
    text-align: center;

    h1 {
      color: $theme-color-1;
      font-size: 2.5rem;
      margin: 0.5rem 0;
    }

    p {
      color: inherit;
      font-size: 1.1rem;
      margin: 0.5rem;
    }
  }

  .card-container {
    display: flex;
    justify-content: center;
    gap: 1.5rem;
    padding: 0 2rem;
    flex-wrap: wrap;
  }

  .card {
    background: linear-gradient(to right, $bg-color-2, $bg-color-1);
    border: 1px solid $bg-color-2;
    border-radius: $border-radius-1;
    padding: 0.9rem;
    color: $txt-color-1;
    text-align: left;
    flex-grow: 1;
    min-width: 10rem;

    h2 {
      font-size: 1.5rem;
      margin: 0 0 0.5rem 0;
    }

    b,
    p {
      font-size: 1rem;
      color: $txt-color-1;
    }

    div.button-container {
      margin-top: 1rem;
      width: 100%;
    }
  }
</style>
