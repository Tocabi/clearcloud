import { render } from '../../__testutils__/render';
import { ManageLibrariesModal } from './ManageLibrariesModal';
import fetchMock from 'fetch-mock';
import { fireEvent, screen, waitFor } from '@testing-library/react';
import { checkPendingMocks } from '../../__testutils__/checkPendingMocks';
import userEvent from '@testing-library/user-event';
import { vi } from 'vitest';

describe('ManageLibrariesModal', () => {
  afterEach(checkPendingMocks);

  describe('Create mode', () => {
    it('can browse through folders', async () => {
      const onClose = vi.fn();
      fetchMock.get(
        'end:/api/libraries/?limit=100',
        {
          totalElements: 0,
          elements: [],
        },
        { overwriteRoutes: true },
      );
      fetchMock.post(
        {
          url: 'end:/api/system/browse',
          body: {
            path: '',
          },
        },
        {
          path: '',
          parent: '',
          folders: [
            {
              path: '/',
            },
          ],
        },
      );

      fetchMock.post(
        {
          url: 'end:/api/system/browse',
          body: {
            path: '/',
          },
          overwriteRoutes: false,
        },
        {
          path: '/',
          parent: '',
          folders: [
            {
              path: '/one',
            },
            {
              path: '/two',
            },
            {
              path: '/three',
            },
          ],
        },
      );
      fetchMock.post(
        {
          url: 'end:/api/system/browse',
          body: {
            path: '/two',
          },
          overwriteRoutes: false,
        },
        {
          path: '/two',
          parent: '/',
          folders: [
            {
              path: '/two/four',
            },
          ],
        },
      );
      await render(<ManageLibrariesModal onClose={onClose} />, {
        loggedIn: true,
      });
      // go into add mode
      await userEvent.click(await screen.findByLabelText('Add Library'));
      // navigate to the root folder
      await waitFor(() => {
        screen.getByText('/');
      });
      await userEvent.click(await screen.findByText('/'));
      await waitFor(() => {
        screen.getByText('/two');
      });
      await userEvent.click(await screen.findByText('/two'));
      // there should be a four now
      await screen.findByText('/two/four');
      // navigate up
      await userEvent.click(screen.getByLabelText('Parent directory'));
      await waitFor(() =>
        expect(
          screen.getByLabelText('Folder:', { selector: 'input' }),
        ).toHaveValue('/'),
      );
    });
  });

  describe('Edit mode', () => {
    it('displays the root folder', async () => {
      const onClose = vi.fn();
      fetchMock.get(
        'end:/api/libraries/?limit=100',
        {
          totalElements: 1,
          elements: [
            { id: 3252, type: 'generic', name: 'Downloads', canWrite: true },
          ],
        },
        { overwriteRoutes: true },
      );
      fetchMock.get('end:/api/libraries/3252', {
        id: 3252,
        name: 'Downloads',
        rootFolder: '/Users/chappio/Downloads',
        sharedWith: [],
        type: 'generic',
      });
      await render(<ManageLibrariesModal onClose={onClose} />, {
        loggedIn: true,
      });
      const input = await screen.findByLabelText('Folder:');
      await waitFor(() => {
        expect(input).toHaveValue('/Users/chappio/Downloads');
      });

      // the first call was loading the libs in the modal
      expect(fetchMock.calls('end:/api/libraries/?limit=100')).toHaveLength(1);

      // now we close the modal which should cause a reload
      fireEvent.click(screen.getByLabelText('Close'));
      await waitFor(() => {
        expect(fetchMock.calls('end:/api/libraries/?limit=100')).toHaveLength(
          2,
        );
      });
    });
  });
});
