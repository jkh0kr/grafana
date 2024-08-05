
import html2canvas from 'html2canvas';
import { useCallback, useEffect, useState } from 'react';

import { openai } from '@grafana/llm';
import { Button } from '@grafana/ui';

export const Hoverbot = () => {
  const [enabled, setEnabled] = useState(false);
  const [loading, setLoading] = useState(false);
  const [reply, setReply] = useState('');

  useEffect(() => {
    openai.enabled().then(setEnabled);
  }, []);

  const ask = useCallback((image: string) => {
    // Stream the completions. Each element is the next stream chunk.
    const stream = openai
      .streamChatCompletions({
        model: openai.Model.LARGE,
        messages: [
          { role: 'system', content: 'You are helping an observability user understand the data they are seeing.' },
          // @ts-expect-error
          { role: 'user', content: [
            {
              "type": "text",
              "text": "Help me understand the following observability data from logs:"
            },
            {
              "type": "image_url",
              "image_url": {
                "url": image
              }
            }
          ]
        }],
      })
      .pipe(
        openai.accumulateContent(),
      );
    // Subscribe to the stream and update the state for each returned value.
    stream.subscribe({
      next: setReply,
      complete: () => {
        setLoading(false);
      },
      error: (e) => {
        console.error(e);
        setLoading(false);
      }
    });
  }, [])

  const helpMe = useCallback(() => {
    if (!enabled) {
      console.error('LLM Disabled');
      return;
    }

    setLoading(true);

    html2canvas(document.querySelector('section')!).then((canvas) => {
      ask(canvas.toDataURL());
    });
  }, [ask, enabled]);

  return (
    <div>
      <Button type="submit" onClick={helpMe} disabled={loading}>
        Help me
      </Button>
      {reply !== '' && (
        <p>{reply}</p>
      )}
    </div>
  );
}
