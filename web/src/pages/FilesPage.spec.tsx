import { checkPendingMocks } from '../__testutils__/checkPendingMocks';
import fetchMock from 'fetch-mock';
import { ReallyDeepPartial, render } from '../__testutils__/render';
import { FilesPage } from './FilesPage';
import { fireEvent, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { RootState } from '../store';
import { formDataMatcher } from '../__testutils__/formDataMatcher';

const initialState: ReallyDeepPartial<RootState> = {
  libraries: {
    libraries: [
      {
        id: 4,
        name: 'Mock Library',
        type: 'generic',
        canWrite: true,
      },
    ],
  },
};

describe('FilesPage', () => {
  afterEach(checkPendingMocks);

  beforeEach(() => {
    fetchMock.getOnce(
      {
        url: 'path:/api/libraries/4/entries',
        query: {
          parent: '',
        },
      },
      {
        status: 200,
        body: [
          {
            name: 'Ballmers Peak Label.xcf',
            parent: '/',
            modified: '2021-03-26T23:32:42.139992387+01:00',
            category: 'Binary',
            size: 2785246,
          },
          {
            name: 'Documents',
            parent: '/',
            modified: '2021-07-01T19:35:16.658563977+02:00',
            category: 'Folder',
            size: 4096,
          },
        ],
      }
    );
  });

  it('shows the files', async () => {
    await render(<FilesPage />, {
      path: '/files/4',
      route: '/files/:library/:path*',
      loggedIn: true,
      initialState,
    });

    expect(await screen.findByText('Mock Library', { selector: 'h1' })).toBeDefined();
    expect(await screen.findByText('Ballmers Peak Label.xcf')).toBeDefined();
    expect(await screen.findByText('2.7 MB')).toBeDefined();
  });

  it('shows file details', async () => {
    await render(<FilesPage />, {
      path: '/files/4',
      route: '/files/:library/:path*',
      loggedIn: true,
      initialState,
    });

    expect(await screen.findByText('Ballmers Peak Label.xcf')).toBeDefined();
    expect(screen.queryByText('Ballmers Peak Label.xcf', { selector: '.details *' })).toBeNull();
    await userEvent.click(screen.getByText('Ballmers Peak Label.xcf'));
    expect(screen.getByText('Ballmers Peak Label.xcf', { selector: '.details *' })).toBeDefined();
    await userEvent.click(screen.getByLabelText('Close details'));
    expect(screen.queryByText('Ballmers Peak Label.xcf', { selector: '.details *' })).toBeNull();
  });

  it('enters a directory on double click', async () => {
    fetchMock.getOnce(
      {
        url: 'path:/api/libraries/4/entries',
        query: {
          parent: 'Documents',
        },
        overwriteRoutes: false,
      },
      {
        status: 200,
        body: [
          {
            name: 'ClearCloud Settings.pdf',
            parent: '/Documents',
            modified: '2021-06-16T09:52:47.769779842+02:00',
            category: 'Binary',
            size: 3570049,
          },
        ],
      }
    );

    const { navigation } = await render(<FilesPage />, {
      path: '/files/4',
      route: '/files/:library/:path*',
      loggedIn: true,
      initialState,
    });

    await userEvent.dblClick(await screen.findByText('Documents'));
    expect(await screen.findByText('ClearCloud Settings.pdf')).toBeDefined();
    expect(navigation.pathname).toBe('/files/4/Documents');
  });

  it('enters the parent directory when going up', async () => {
    fetchMock.getOnce(
      {
        url: 'path:/api/libraries/4/entries',
        query: {
          parent: 'Documents',
        },
        overwriteRoutes: false,
      },
      {
        status: 200,
        body: [
          {
            name: 'ClearCloud Settings.pdf',
            parent: '/Documents',
            modified: '2021-06-16T09:52:47.769779842+02:00',
            category: 'Binary',
            size: 3570049,
          },
        ],
      }
    );

    const { navigation } = await render(<FilesPage />, {
      path: '/files/4/Documents',
      route: '/files/:library/:path*',
      loggedIn: true,
      initialState,
    });

    expect(await screen.findByText('ClearCloud Settings.pdf')).toBeDefined();
    await userEvent.click(screen.getByLabelText('Parent directory'));
    expect(await screen.findByText('Documents')).toBeDefined();
    expect(navigation.pathname).toBe('/files/4');
  });

  it('uploads files', async () => {
    const fileOne = new File(['file content here'], 'upload.txt');
    const fileOneEntry = {
      name: 'upload.txt',
      parent: '/',
      modified: '2021-03-26T23:32:42.139992387+01:00',
      category: 'Document',
      size: 17,
    };
    const fileTwo = new File(['another file? in this economy?'], 'another.zip');
    const fileTwoEntry = {
      name: 'another.zip',
      parent: '/',
      modified: '2021-03-26T23:32:42.139992387+01:00',
      category: 'Archive',
      size: 1337,
    };

    fetchMock.postOnce(
      {
        url: 'path:/api/libraries/4/entries',
        matcher: formDataMatcher({
          name: 'upload.txt',
          parent: '',
          data: fileOne,
        }),
      },
      {
        status: 201,
        body: fileOneEntry,
      }
    );
    fetchMock.getOnce(
      {
        url: 'path:/api/libraries/4/entries',
        query: {
          parent: '',
        },
        overwriteRoutes: false,
      },
      {
        status: 200,
        body: [fileOneEntry],
      }
    );
    fetchMock.postOnce(
      {
        url: 'path:/api/libraries/4/entries',
        matcher: formDataMatcher({
          name: 'another.zip',
          parent: '',
          data: fileTwo,
        }),
        overwriteRoutes: false,
      },
      {
        status: 201,
        body: fileTwoEntry,
      }
    );
    fetchMock.getOnce(
      {
        url: 'path:/api/libraries/4/entries',
        query: {
          parent: '',
        },
        overwriteRoutes: false,
      },
      {
        status: 200,
        body: [fileOneEntry, fileTwoEntry],
      }
    );

    await render(<FilesPage />, {
      path: '/files/4',
      route: '/files/:library/:path*',
      loggedIn: true,
      initialState,
    });

    fireEvent.change(screen.getByTestId('file-input'), {
      target: {
        files: [fileOne, fileTwo],
      },
    });
    await screen.findByText('upload.txt');
    await screen.findByText('another.zip');
  });
});
