# DB Creation SQL

```sql
-- Image table.
CREATE TABLE "images" (
  "id"                    BIGSERIAL,
  "telegram_user_id"      BIGINT       NOT NULL, -- ID for the user sending the image, for validating image deleting callbacks
  "telegram_reply_msg_id" BIGINT       NOT NULL, -- ID for the replying message with the image URL, for updating the reply
  "imgur_url"             VARCHAR(100) NOT NULL, -- URL of the image uploaded to Imgur
  "imgur_delete_hash"     VARCHAR(20)  NOT NULL, -- Hash for deleting the image from Imgur
  "create_time_utc"       TIMESTAMP    NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id)
);

CREATE UNIQUE INDEX uniq_idx_images_telegram_reply_msg_id ON images ("telegram_reply_msg_id") WHERE "telegram_reply_msg_id" != 0;
CREATE UNIQUE INDEX uniq_idx_images_imgur_url ON images ("imgur_url");
CREATE UNIQUE INDEX uniq_idx_images_imgur_delete_hash ON images ("imgur_delete_hash");
```
