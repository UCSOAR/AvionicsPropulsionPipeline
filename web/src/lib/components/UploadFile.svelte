<script lang="ts">
  import { UploadCloud } from "lucide-svelte";
  import { endpointMapping } from "$lib/utils/constants";
  import IconButton from "$lib/components/IconButton.svelte";

  // Props
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
        // Make sure the URL is correct
        const response = await fetch(endpointMapping.uploadStaticFireUrl, {
          method: "POST",
          body: formData,
          // Adding these headers might help with certain server configurations
          // headers: {
          //   'Accept': 'application/json',
          // },
        });

        if (response.ok) {
          console.log(`File ${file.name} uploaded successfully`);
        } else {
          const errorText = await response.text();
          console.error(`Failed to upload file ${file.name}: ${errorText}`);
          // You might want to show an error message to the user here
        }
      } catch (error) {
        console.error(`Error uploading file ${file.name}:`, error);
        // You might want to show an error message to the user here
      }
    }

    uploading = false;
    target.value = ''; // Reset the file input
    
    // Only call onUploadComplete if at least one file was successfully uploaded
    onUploadComplete();
  };

  const handleClick = () => {
    console.log("I reached here");
    if (fileInput) {
      fileInput.click();
    }
  };
</script>

<div class="uploader-container">
  <!-- Hidden file input -->
  <input
    type="file"
    bind:this={fileInput}
    onchange={handleFileChange}
    hidden
    accept=".lvm,"
  />

  <!-- Upload Button -->
  <IconButton
    class="upload-btn"
    icon={UploadCloud}
    label={uploading ? "Uploading..." : "Upload File"}
    onclick={handleClick}
    disabled={uploading}
  />
</div>

<style lang="scss">
  .uploader-container {
    width: 100%;
  }
  
  .upload-btn{
    width:15rem;
  }
</style>