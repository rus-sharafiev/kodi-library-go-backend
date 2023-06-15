SELECT jsonb_agg(r) result
FROM (

    SELECT u.*, json_agg(pr) profile
    FROM public."User" u
    LEFT JOIN (

        SELECT p.*, json_agg(sr) specialties
        FROM public."Profile" p
        LEFT JOIN (
            SELECT s.*,
                json_agg(o) offers
            FROM public."Specialty" s
                LEFT JOIN public."Offer" o ON s."profileUserId" = o."specialtyProfileUserId"
            GROUP BY s."subcategoryId",
                s."profileUserId"
        ) sr 
        ON p."userId" = sr."profileUserId"
        GROUP BY p."userId"

    ) pr 
    ON u.id = pr."userId"
    where id = 1
    GROUP BY u.id
) r;