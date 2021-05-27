import React, { useState } from 'react';
import { SettingsLayout } from '../../components/layout/SettingsLayout';
import { Form } from '../../components/input/Form';
import { TextInput } from '../../components/input/TextInput';

export const NewUserPage: React.FC = () => {
  const [username, setUsername] = useState('');
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');

  return (
    <SettingsLayout className="new-user-page">
      <h2 className="title">New User</h2>
      <Form onSubmit={() => undefined} submitLabel="Create">
        <div className="columns">
          <div>
            <label htmlFor="username">Username</label>
            <label htmlFor="firstName">First Name</label>
            <label htmlFor="lastName">Last Name</label>
          </div>
          <div className="inputs">
            <TextInput id="username" value={username} onChange={setUsername} placeholder="Username" />
            <TextInput id="firstName" value={lastName} onChange={setLastName} placeholder="First Name" />
            <TextInput id="lastName" value={firstName} onChange={setFirstName} placeholder="Last Name" />
          </div>
        </div>
      </Form>
    </SettingsLayout>
  );
};
