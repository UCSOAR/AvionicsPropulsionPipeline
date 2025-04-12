<script lang="ts">
  import IconButton from "$lib/components/IconButton.svelte";
  import { UploadCloud } from "@lucide/svelte";
  import { endpointMapping } from "$lib/utils/constants";

  export let onUploadComplete: () => void = () => {};

  let fileInput: HTMLInputElement;
  let uploading = false;

  const handleFileChange = async (event: Event) => {
    const target = event.target as HTMLInputElement;
    if (!target.files || target.files.length === 0) return;

    const files = Array.from(target.files);

    uploading = true;

    for (const file of files) {
      const formData = new FormData();
      formData.append("file", file);

      try {
        const response = await fetch(endpointMapping.uploadStaticFireUrl, {
          method: "POST",
          body: formData,
        });

        if (response.ok) {
          console.log(`File ${file.name} uploaded successfully`);
        } else {
          const errorText = await response.text();
          console.error(`Failed to upload file ${file.name}: ${errorText}`);
        }
      } catch (error) {
        console.error(`Error uploading file ${file.name}:`, error);
      }
    }

    uploading = false;
    target.value = "";

    onUploadComplete();
  };

  const handleClick = () => {
    if (fileInput) {
      fileInput.click();
    }
  };
</script>

<div class="uploader-container">
  <input
    type="file"
    bind:this={fileInput}
    onchange={handleFileChange}
    hidden
    accept=".lvm,"
  />
  <div class="upload-btn">
    <IconButton
      icon={UploadCloud}
      label={uploading ? "Uploading..." : "Upload File"}
      onClick={handleClick}
      isDisabled={uploading}
    />
  </div>
</div>

<style lang="scss">
  .uploader-container {
    width: 100%;
    display: flex;
    justify-content: center;
  }
</style>
