import { css } from '@emotion/css';
import {
  memo,
  useState,
  // ChangeEvent
} from 'react';

import {
  // VariableSuggestion,
  GrafanaTheme2,
  // DataLink,
} from '@grafana/data';

// TODO: Move this to @grafana/data
/**
 * Calculation configuration.
 */
export interface Calculation {
  name: string;
  expression: string;
}

import { useStyles2 } from '../../themes/index';
// import { isCompactUrl } from '../../utils/dataLinks';
import { Field } from '../Forms/Field';
import { Input } from '../Input/Input';

// import { DataLinkInput } from './DataLinkInput';

interface CalculationEditorProps {
  // index: number;
  // isLast: boolean;
  value: Calculation;
  // suggestions: VariableSuggestion[];
  // onChange: (index: number, link: DataLink, callback?: () => void) => void;
}

const getStyles = (theme: GrafanaTheme2) => ({
  listItem: css({
    marginBottom: theme.spacing(),
  }),
  infoText: css({
    paddingBottom: theme.spacing(2),
    // marginLeft: '66px',
    color: theme.colors.text.secondary,
  }),
});

export const CalculationEditor = memo(
  ({
    //  index,
    value,
    // onChange, suggestions, isLast
  }: CalculationEditorProps) => {
    const [name, setName] = useState('');
    const [calculation, setCalculation] = useState('');

    const styles = useStyles2(getStyles);

    // const handleNameChange = (event: ChangeEvent<HTMLInputElement>) => {
    //   setName(event.target.value);
    // };

    // const onUrlChange = (url: string, callback?: () => void) => {
    //   onChange(index, { ...value, url }, callback);
    // };
    // const onTitleChange = (event: ChangeEvent<HTMLInputElement>) => {
    //   // onChange(index, { ...value, title: event.target.value });
    // };

    return (
      <div className={styles.listItem}>
        <Field label="Name">
          <Input
            value={name}
            onChange={(event) => setName(event.currentTarget.value)}
            placeholder="Custom calculation"
          />
        </Field>

        <Field
          label="Calculation"
          // invalid={isCompactUrl(value.url)}
          // error="Data link is an Explore URL in a deprecated format. Please visit the URL to be redirected, and edit this data link to use that URL."
        >
          {/* <DataLinkInput value={value.url} onChange={onUrlChange} suggestions={suggestions} /> */}
          <Input
            value={calculation}
            onChange={(event) => setCalculation(event.currentTarget.value)}
            placeholder="7 + 4 / 2"
          />
        </Field>

        <div className={styles.infoText}>
          You can reference data variables like series name, labels and values. Type CMD+Space, CTRL+Space, or $ to open
          variable suggestions.
        </div>
      </div>
    );
  }
);

CalculationEditor.displayName = 'CalculationEditor';
