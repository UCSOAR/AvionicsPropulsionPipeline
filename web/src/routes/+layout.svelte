<script lang="ts">
  import { userStore } from "$lib/stores/userStore";
  import { fetchMe } from "$lib/utils/getMe";
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";

  const { children } = $props();

  onMount(() => {
    (async () => {
      const me = await fetchMe();

      if (!me) {
        await goto("/");
        return;
      }

      $userStore = me;
    })();
  });
</script>

{@render children()}
