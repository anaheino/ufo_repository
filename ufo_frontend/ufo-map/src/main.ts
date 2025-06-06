import './assets/main.css';

import { createApp } from 'vue';
import { createPinia } from 'pinia';

import '@mdi/font/css/materialdesignicons.css';
import App from './App.vue';
import router from './router';
import 'vuetify/styles';
import { createVuetify } from 'vuetify';
import * as vuetifyComponents from 'vuetify/components';
import * as directives from 'vuetify/directives';

const vuetify = createVuetify({
    components: {
        ...vuetifyComponents,
    },
    directives,
});

createApp(App)
    .use(createPinia())
    .use(router)
    .use(vuetify)
    .mount('#app');
