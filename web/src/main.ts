import { createApp } from "vue";
import App from "./App.vue";
import VueApexCharts from 'vue3-apexcharts';
import HighchartsVue from 'highcharts-vue';
import { createPinia } from 'pinia';


const app = createApp(App);
app.use(VueApexCharts);
app.use(HighchartsVue)
const pinia = createPinia();
app.use(pinia);
app.component('apexchart', VueApexCharts);
app.mount('#app');  