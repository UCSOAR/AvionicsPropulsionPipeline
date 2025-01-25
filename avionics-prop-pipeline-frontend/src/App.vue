<script setup lang="ts">
import './styles/app.scss';

import BucketUploadForm from './components/BucketUploadForm.vue';
import BucketUploads from './components/BucketUploads.vue';
import Header from './components/Header.vue';
</script>

<template>
  <Header />
  <main>
    <h2>Upload a file</h2>
    <BucketUploadForm />

    <h2>Uploaded files</h2>
    <BucketUploads />
  </main>
</template>

<style scoped></style>
