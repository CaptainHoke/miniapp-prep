with
paying_users_per_date as
(
    select
        time::date as date,
        count(distinct user_id) filter (
            where order_id not in (
                select order_id from user_actions where action = 'cancel_order'
            )
        ) as paying_users
    from
        user_actions
    group by date
),
active_couriers_per_date as
(
    select
        time::date as date,
        count(distinct courier_id) filter (
            where order_id not in (
                select order_id from user_actions where action = 'cancel_order'
            )
        ) as active_couriers
    from
        courier_actions
    group by date
),
platform_users_per_date as
(
    select
        first_action_date as date,
        sum(new_users) over (order by first_action_date)::integer as total_users,
        sum(new_couriers) over (order by first_action_date)::integer as total_couriers
    from
    (
        select
            first_action_date,
            count(user_id) as new_users
        from
        (
            select
                user_id,
                min(time::date) as first_action_date
            from user_actions
            group by user_id
        ) t1
        group by first_action_date
    ) t2
    inner join
    (
        select
            first_action_date,
            count(courier_id) as new_couriers
        from
        (
            select
                courier_id,
                min(time::date) as first_action_date
            from courier_actions
            group by courier_id
        ) t3
        group by first_action_date
    ) t4
    using (first_action_date)
)

select
    date,
    paying_users,
    active_couriers,
    round(paying_users * 100.0 / total_users, 2) as paying_users_share,
    round(active_couriers * 100.0 / total_couriers, 2) as active_couriers_share
from
    platform_users_per_date
inner join
    paying_users_per_date
using (date)
inner join
    active_couriers_per_date
using (date)
