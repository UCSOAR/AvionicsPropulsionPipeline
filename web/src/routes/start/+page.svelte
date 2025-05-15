<script lang="ts">
  import "$lib/styles/global.scss";
  import IconButton from "$lib/components/IconButton.svelte";
  import SplashHeader from "$lib/components/SplashHeader.svelte";
  import AccountInfo from "$lib/components/AccountInfo.svelte";
  import { goto } from "$app/navigation";
  import { LucideX, Rocket } from "@lucide/svelte";

  const gotoTAC = () => {
    goto("/tac");
  };
</script>

<main class="splash-main">
  <div class="account-info">
    <AccountInfo />
  </div>

  <div class="splash-content-container">
    <SplashHeader />
    <div class="card-container">
      <div class="card">
        <h2>TAC</h2>
        <p>
          <b>Test Analysis Code</b> developed for the purpose of analysing data obtained
          from hybrid engine tests.
        </p>
        <div class="button-container">
          <IconButton icon={Rocket} label="Get Started" onClick={gotoTAC} />
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

<style lang="scss" scoped>
  div.account-info {
    padding: 0.9rem;
  }
</style>
