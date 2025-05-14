<script lang="ts">
  import "$lib/styles/global.scss";
  import IconButton from "$lib/components/IconButton.svelte";
  import SplashHeader from "$lib/components/SplashHeader.svelte";
  import { KeyIcon } from "@lucide/svelte";
  import { endpointMapping, redirectUriParam } from "$lib/utils/constants";

  const gotoGoogleLogin = () => {
    const redirectUri = window.location.href;
    const endpoint = new URL(endpointMapping.getGoogleLoginUrl);
    endpoint.searchParams.set(redirectUriParam, redirectUri);

    window.location.href = endpoint.toString();
  };
</script>

<main class="splash-main">
  <div class="splash-content-container">
    <SplashHeader />
    <div class="card-container">
      <div class="card">
        <h2>Access Restricted</h2>
        <p>
          Please authenticate with your <b>Google</b> account credentials below
          to access the pipeline. In order to use the pipeline, you must be
          entered into the <b>whitelist</b>.
        </p>
        <div class="button-container">
          <IconButton
            icon={KeyIcon}
            label="Authenticate With Google"
            onClick={gotoGoogleLogin}
          />
        </div>
      </div>
    </div>
  </div>
</main>
