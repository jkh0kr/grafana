import { css } from '@emotion/css';
import { DOMAttributes } from '@react-types/shared';
import { memo, forwardRef, useCallback } from 'react';
import { useLocation } from 'react-router-dom';

import { GrafanaTheme2, NavModelItem } from '@grafana/data';
import { selectors } from '@grafana/e2e-selectors';
import { config, reportInteraction, i18n } from '@grafana/runtime';
import { CustomScrollbar, Icon, IconButton, useStyles2, Stack } from '@grafana/ui';
import { useGrafana } from 'app/core/context/GrafanaContext';
import { setBookmark } from 'app/core/reducers/navBarTree';
import { usePatchUserPreferencesMutation } from 'app/features/preferences/api/index';
import { useDispatch, useSelector } from 'app/types';

import { MegaMenuItem } from './MegaMenuItem';
import { usePinnedItems } from './hooks';
import { enrichWithInteractionTracking, getActiveItem } from './utils';

export const MENU_WIDTH = '300px';

export interface Props extends DOMAttributes {
  onClose: () => void;
}

export const MegaMenu = memo(
  forwardRef<HTMLDivElement, Props>(({ onClose, ...restProps }, ref) => {
    const navTree = useSelector((state) => state.navBarTree);
    const styles = useStyles2(getStyles);
    const location = useLocation();
    const { chrome } = useGrafana();
    const dispatch = useDispatch();
    const state = chrome.useState();
    const [patchPreferences] = usePatchUserPreferencesMutation();
    const pinnedItems = usePinnedItems();

    // Remove profile + help from tree
    const navItems = navTree
      .filter((item) => item.id !== 'profile' && item.id !== 'help')
      .map((item) => enrichWithInteractionTracking(item, state.megaMenuDocked));

    const activeItem = getActiveItem(navItems, state.sectionNav.node, location.pathname);

    const handleDockedMenu = () => {
      chrome.setMegaMenuDocked(!state.megaMenuDocked);
      if (state.megaMenuDocked) {
        chrome.setMegaMenuOpen(false);
      }

      // refocus on undock/menu open button when changing state
      setTimeout(() => {
        document.getElementById(state.megaMenuDocked ? 'mega-menu-toggle' : 'dock-menu-button')?.focus();
      });
    };

    const isPinned = useCallback(
      (url?: string) => {
        if (!url || !pinnedItems?.length) {
          return false;
        }
        return pinnedItems?.includes(url);
      },
      [pinnedItems]
    );

    const onPinItem = (item: NavModelItem) => {
      const url = item.url;
      if (url && config.featureToggles.pinNavItems) {
        const isSaved = isPinned(url);
        const newItems = isSaved ? pinnedItems.filter((i) => url !== i) : [...pinnedItems, url];
        const interactionName = isSaved ? 'grafana_nav_item_unpinned' : 'grafana_nav_item_pinned';
        reportInteraction(interactionName, {
          path: url,
        });
        patchPreferences({
          patchPrefsCmd: {
            navbar: {
              bookmarkUrls: newItems,
            },
          },
        }).then((data) => {
          if (!data.error) {
            dispatch(setBookmark({ item: item, isSaved: !isSaved }));
          }
        });
      }
    };

    return (
      <div data-testid={selectors.components.NavMenu.Menu} ref={ref} {...restProps}>
        <div className={styles.mobileHeader}>
          <Icon name="bars" size="xl" />
          <IconButton
            tooltip={i18n.t('navigation.megamenu.close', 'Close menu')}
            name="times"
            onClick={onClose}
            size="xl"
            variant="secondary"
          />
        </div>
        <nav className={styles.content}>
          <CustomScrollbar showScrollIndicators hideHorizontalTrack>
            <ul className={styles.itemList} aria-label={i18n.t('navigation.megamenu.list-label', 'Navigation')}>
              {navItems.map((link, index) => (
                <Stack key={link.text} direction={index === 0 ? 'row-reverse' : 'row'} alignItems="center">
                  {index === 0 && (
                    <IconButton
                      id="dock-menu-button"
                      className={styles.dockMenuButton}
                      tooltip={
                        state.megaMenuDocked
                          ? i18n.t('navigation.megamenu.undock', 'Undock menu')
                          : i18n.t('navigation.megamenu.dock', 'Dock menu')
                      }
                      name="web-section-alt"
                      onClick={handleDockedMenu}
                      variant="secondary"
                    />
                  )}
                  <MegaMenuItem
                    link={link}
                    isPinned={isPinned}
                    onClick={state.megaMenuDocked ? undefined : onClose}
                    activeItem={activeItem}
                    onPin={onPinItem}
                  />
                </Stack>
              ))}
            </ul>
          </CustomScrollbar>
        </nav>
      </div>
    );
  })
);

MegaMenu.displayName = 'MegaMenu';

const getStyles = (theme: GrafanaTheme2) => ({
  content: css({
    display: 'flex',
    flexDirection: 'column',
    height: '100%',
    minHeight: 0,
    position: 'relative',
  }),
  mobileHeader: css({
    display: 'flex',
    justifyContent: 'space-between',
    padding: theme.spacing(1, 1, 1, 2),
    borderBottom: `1px solid ${theme.colors.border.weak}`,

    [theme.breakpoints.up('md')]: {
      display: 'none',
    },
  }),
  itemList: css({
    boxSizing: 'border-box',
    display: 'flex',
    flexDirection: 'column',
    listStyleType: 'none',
    padding: theme.spacing(1, 1, 2, 1),
    [theme.breakpoints.up('md')]: {
      width: MENU_WIDTH,
    },
  }),
  dockMenuButton: css({
    display: 'none',

    [theme.breakpoints.up('xl')]: {
      display: 'inline-flex',
    },
  }),
});
