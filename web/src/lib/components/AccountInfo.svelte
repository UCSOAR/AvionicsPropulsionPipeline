<script lang="ts">
  import IconButton from "./IconButton.svelte";
  import { userStore } from "$lib/stores/userStore";
  import { LogOut } from "@lucide/svelte";
  import { endpointMapping } from "$lib/utils/constants";
  import { goto } from "$app/navigation";

  const handleLogout = async () => {
    await fetch(endpointMapping.postLogoutUrl, {
      method: "POST",
      credentials: "include",
    });

    $userStore = null;
    await goto("/");
  };
</script>

{#if $userStore}
  <div class="container">
    <img alt="Account Profile" src={$userStore.picture} />
    <IconButton
      label={`Logout ${$userStore.email}`}
      icon={LogOut}
      onClick={handleLogout}
    />
  </div>
{/if}

<style lang="scss" scoped>
  @use "../styles/variables.scss" as *;

  div.container {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: row;
    gap: 0.5rem;

    img {
      border-radius: $border-radius-1;
      border: 1px solid $outline-color-1;
      height: 2.4rem;
      width: auto;
    }
  }
</style>
