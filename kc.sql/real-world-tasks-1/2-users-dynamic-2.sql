select
    date,
    new_users,
    new_couriers,
    total_users,
    total_couriers,
    new_users_change,
    new_couriers_change,
    round(new_users * 100.0 / lag(total_users, 1) over (), 2) as total_users_growth,
    round(new_couriers * 100.0 / lag(total_couriers, 1) over (), 2) as total_couriers_growth
from
(
    select
        first_action_date as date,
        new_users,
        new_couriers,
        sum(new_users) over (order by first_action_date)::integer as total_users,
        sum(new_couriers) over (order by first_action_date)::integer as total_couriers,
        round((new_users - lag(new_users, 1) over ()) * 100.0 / lag(new_users, 1) over (), 2) as new_users_change,
        round((new_couriers - lag(new_couriers, 1) over ()) * 100.0 / lag(new_couriers, 1) over (), 2) as new_couriers_change
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
) t5
