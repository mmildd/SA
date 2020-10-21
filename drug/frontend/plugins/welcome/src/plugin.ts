import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Tables from './components/Table';
import logins from './components/Login';
import home from './components/home';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', logins);
    router.registerRoute('/home', home);
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/Tables', Tables);
  },
});
