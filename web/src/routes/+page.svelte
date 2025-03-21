<script lang="ts">
  import "../styles/global.scss";
  import "$lib/Sidebar.svelte";
  import IconButton from "$lib/Sidebar.svelte";
  import { UploadCloudIcon } from "@lucide/svelte";
</script>

<IconButton icon={UploadCloudIcon} label="Upload File" />
