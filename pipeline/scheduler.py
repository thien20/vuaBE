import logging
import pytz
from apscheduler.schedulers.blocking import BlockingScheduler
from utils.daily_consumer_utils import trigger_scrape

logging.basicConfig(level=logging.DEBUG, format='%(asctime)s - %(levelname)s - %(message)s')
logger = logging.getLogger(__name__)

if __name__ == '__main__':
    scheduler = BlockingScheduler(timezone=pytz.timezone('Asia/Ho_Chi_Minh'))
    logger.info("Scheduler initialized with UTC timezone")

    scheduler.add_job(
        trigger_scrape, 'cron', day='*', hour=8,
        args=['scrape']
    )

    try:
        logger.info("Scheduler started. Waiting for scheduled jobs...")
        scheduler.start()
    except (KeyboardInterrupt, SystemExit):
        logger.info("Scheduler stopped.")
