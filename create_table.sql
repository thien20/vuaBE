-- CREATE TABLE thoi_su (ID INT, link TEXT, title TEXT, content TEXT);
-- CREATE TABLE the_gioi (ID INT, link TEXT, title TEXT, content TEXT);
-- CREATE TABLE kinh_doanh (ID INT, link TEXT, title TEXT, content TEXT);
-- CREATE TABLE giai_tri (ID INT, link TEXT, title TEXT, content TEXT);
-- CREATE TABLE the_thao (ID INT, link TEXT, title TEXT, content TEXT);
-- CREATE TABLE phap_luat (ID INT, link TEXT, title TEXT, content TEXT);
-- CREATE TABLE giao_duc (ID INT, link TEXT, title TEXT, content TEXT);
-- CREATE TABLE suc_khoe (ID INT, link TEXT, title TEXT, content TEXT);
-- CREATE TABLE doi_song (ID INT, link TEXT, title TEXT, content TEXT);
-- CREATE TABLE du_lich (ID INT, link TEXT, title TEXT, content TEXT);
-- CREATE TABLE khoa_hoc (ID INT, link TEXT, title TEXT, content TEXT);
-- CREATE TABLE so_hoa (ID INT, link TEXT, title TEXT, content TEXT);
-- CREATE TABLE oto_xemay (ID INT, link TEXT, title TEXT, content TEXT);
-- CREATE TABLE y_kien (ID INT, link TEXT, title TEXT, content TEXT);
-- CREATE TABLE tam_su (ID INT, link TEXT, title TEXT, content TEXT);

-- CREATE TABLE IF NOT EXISTS thoi_su (ID INT, link TEXT, title TEXT, content TEXT);
-- CREATE TABLE IF NOT EXISTS the_gioi (ID INT, link TEXT, title TEXT, content TEXT);
-- CREATE TABLE IF NOT EXISTS goc_nhin (ID INT, link TEXT, title TEXT, content TEXT);

-- gorm -> lower(snakes)
CREATE TABLE IF NOT EXISTS news (id INT, link TEXT, title TEXT, content TEXT, category TEXT);
CREATE TABLE IF NOT EXISTS category (category_id INT, category_name TEXT);
CREATE TABLE IF NOT EXISTS news_categories (category_id INT, news_id INT);