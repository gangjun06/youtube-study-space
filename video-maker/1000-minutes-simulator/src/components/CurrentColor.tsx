import { css } from '@emotion/react'
import { FC } from 'react'
import { MdInvertColors } from 'react-icons/md'
import { Rank, ranks } from '../lib/ranks'
import * as styles from '../styles/CurrentColor.styles'
import * as common from '../styles/common.styles'

type Props = {
    elapsedMinutes: number
}

const CurrentColor: FC<Props> = (props) => {
    let currentColorCode = 'inherit'
    ranks.forEach((rank: Rank) => {
        if (
            rank.FromHours <= props.elapsedMinutes &&
            props.elapsedMinutes < rank.ToHours
        ) {
            currentColorCode = rank.ColorCode
        }
    })

    return (
        <div css={styles.currentColor}>
            <div css={styles.innerCell}>
                <div css={common.heading}>
                    <MdInvertColors
                        size={common.IconSize}
                        css={styles.icon}
                    ></MdInvertColors>
                    <span>達成カラー</span>
                </div>
                <div
                    css={css`
                        ${styles.colorBox};
                        background-color: ${currentColorCode};
                        box-shadow: 0 0 50px ${currentColorCode};
                    `}
                ></div>
                <div css={styles.annotation}>
                    {/* 値は時間ではなく分を使うことに注意 */}
                    実際に累計作業時間が{props.elapsedMinutes}
                    時間に達すると<br></br>座席がこの色になります。
                </div>
            </div>
        </div>
    )
}

export default CurrentColor
