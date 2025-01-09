<script setup lang="ts">
import BucketUploadForm from './components/BucketUploadForm.vue';
import BucketUploads from './components/BucketUploads.vue';
</script>

<template>
  <main>
    <h1>SOAR Avionics-Propulsion Pipeline</h1>
    <h2>Upload a file</h2>
    <BucketUploadForm />

    <h2>Uploaded files</h2>
    <BucketUploads />
  </main>
</template>

<style scoped></style>
